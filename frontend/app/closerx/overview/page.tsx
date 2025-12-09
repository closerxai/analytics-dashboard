'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { DataTable, Column } from '@/components/analytics/data-table';
import { closerxOverviewData } from '@/lib/mock-data';
import { Users, Megaphone, Phone, Clock, Zap } from 'lucide-react';

export default function CloserXOverview() {
  const columns: Column<typeof closerxOverviewData.topAgencies[0]>[] = [
    { key: 'name', label: 'Agency Name', sortable: true },
    {
      key: 'minutes',
      label: 'Minutes',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
  ];

  return (
    <div className="space-y-8">
      <PageHeader
        title="CloserX Overview"
        description="Campaign management and call analytics"
      />

      <KPIGrid columns={3}>
        <MetricCard
          label="Total Users"
          value={closerxOverviewData.kpis.totalUsers}
          icon={Users}
        />
        <MetricCard
          label="Active Campaigns"
          value={closerxOverviewData.kpis.activeCampaigns}
          icon={Megaphone}
        />
        <MetricCard
          label="Quick Calls"
          value={closerxOverviewData.kpis.quickCalls}
          icon={Phone}
        />
        <MetricCard
          label="Campaign Calls"
          value={closerxOverviewData.kpis.campaignCalls}
          icon={Phone}
        />
        <MetricCard
          label="Total Minutes"
          value={closerxOverviewData.kpis.totalMinutes}
          icon={Clock}
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Credit Burn"
          value={closerxOverviewData.kpis.creditBurn}
          icon={Zap}
          formatValue={(val) => val.toLocaleString()}
        />
      </KPIGrid>

      <div className="grid gap-6 lg:grid-cols-2">
        <LineAreaChart
          title="Calls Per Day"
          data={closerxOverviewData.callsPerDay}
          valueFormatter={(val) => val.toLocaleString()}
        />
        <LineAreaChart
          title="Minutes Per Day"
          data={closerxOverviewData.minutesPerDay}
          valueFormatter={(val) => `${val.toLocaleString()} min`}
        />
      </div>

      <LineAreaChart
        title="Credit Burn Trend"
        data={closerxOverviewData.creditBurnTrend}
        valueFormatter={(val) => val.toLocaleString()}
      />

      <DataTable
        title="Top 5 Agencies by Minutes"
        data={closerxOverviewData.topAgencies}
        columns={columns}
        searchable={false}
        pageSize={5}
      />
    </div>
  );
}
