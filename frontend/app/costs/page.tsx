'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { PieChart } from '@/components/analytics/pie-chart';
import { BarChart } from '@/components/analytics/bar-chart';
import { DataTable, Column } from '@/components/analytics/data-table';
import { costsData } from '@/lib/mock-data';
import { Server, Cloud } from 'lucide-react';
import { Badge } from '@/components/ui/badge';

export default function CostsPage() {
  const columns: Column<typeof costsData.costDetails[0]>[] = [
    { key: 'provider', label: 'Provider', sortable: true },
    { key: 'service', label: 'Service', sortable: true },
    {
      key: 'spend',
      label: 'Spend',
      sortable: true,
      render: (val) => `$${val.toLocaleString()}`,
    },
    {
      key: 'change',
      label: 'Change',
      sortable: true,
      render: (val) => {
        const isPositive = val > 0;
        return (
          <Badge variant={isPositive ? 'destructive' : 'secondary'}>
            {isPositive ? '+' : ''}
            {val}%
          </Badge>
        );
      },
    },
    { key: 'notes', label: 'Notes' },
  ];

  return (
    <div className="space-y-8">
      <PageHeader
        title="Infrastructure Costs"
        description="Third-party service costs and infrastructure spend"
      />

      <KPIGrid columns={4}>
        <MetricCard
          label="AWS Spend"
          value={costsData.kpis.awsCost}
          icon={Cloud}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="GCP Spend"
          value={costsData.kpis.gcpCost}
          icon={Cloud}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Twilio/Ultravox"
          value={costsData.kpis.twilioCost}
          icon={Server}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Total Infra Cost"
          value={costsData.kpis.totalInfraCost}
          icon={Server}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
      </KPIGrid>

      <div className="grid gap-6 lg:grid-cols-2">
        <PieChart
          title="Provider Breakdown"
          data={costsData.providerBreakdown}
          valueFormatter={(val) => `$${val.toLocaleString()}`}
        />
        <LineAreaChart
          title="Cost Trend"
          data={costsData.costTrend}
          valueFormatter={(val) => `$${val.toLocaleString()}`}
        />
      </div>

      <BarChart
        title="Cost by Service"
        data={costsData.perServiceCosts}
        dataKey="cost"
        categoryKey="service"
        valueFormatter={(val) => `$${val.toLocaleString()}`}
      />

      <DataTable
        title="Detailed Cost Breakdown"
        data={costsData.costDetails}
        columns={columns}
        searchPlaceholder="Search by provider or service..."
      />
    </div>
  );
}
