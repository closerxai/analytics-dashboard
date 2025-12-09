'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { PieChart } from '@/components/analytics/pie-chart';
import { DataTable, Column } from '@/components/analytics/data-table';
import { snowieOverviewData } from '@/lib/mock-data';
import { Users, Phone, Clock, Zap } from 'lucide-react';

export default function SnowieOverview() {
  const callsAndMinutesData = snowieOverviewData.callsTrend.map((item) => ({
    date: item.date,
    'Calls': item.calls,
    'Minutes': item.minutes,
  }));

  const columns: Column<typeof snowieOverviewData.dailySummary[0]>[] = [
    {
      key: 'date',
      label: 'Date',
      sortable: true,
      render: (val) => new Date(val).toLocaleDateString(),
    },
    {
      key: 'calls',
      label: 'Calls',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
    {
      key: 'duration',
      label: 'Duration (min)',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
    {
      key: 'credits',
      label: 'Credits',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
  ];

  return (
    <div className="space-y-8">
      <PageHeader
        title="Snowie Overview"
        description="AI model usage and performance analytics"
      />

      <KPIGrid columns={4}>
        <MetricCard
          label="Total Users"
          value={snowieOverviewData.kpis.totalUsers}
          icon={Users}
        />
        <MetricCard
          label="Total Calls"
          value={snowieOverviewData.kpis.totalCalls}
          icon={Phone}
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Total Duration"
          value={snowieOverviewData.kpis.totalDuration}
          icon={Clock}
          suffix=" min"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Total Credits"
          value={snowieOverviewData.kpis.totalCredits}
          icon={Zap}
          formatValue={(val) => val.toLocaleString()}
        />
      </KPIGrid>

      <div className="grid gap-6 lg:grid-cols-2">
        <LineAreaChart
          title="Calls & Minutes Trend"
          data={callsAndMinutesData}
          dataKeys={['Calls', 'Minutes']}
          colors={['hsl(var(--chart-1))', 'hsl(var(--chart-2))']}
          valueFormatter={(val) => val.toLocaleString()}
        />
        <PieChart
          title="Credits by Model"
          data={snowieOverviewData.modelCredits}
          valueFormatter={(val) => val.toLocaleString()}
        />
      </div>

      <LineAreaChart
        title="Credits Usage Trend"
        data={snowieOverviewData.creditsTrend}
        valueFormatter={(val) => val.toLocaleString()}
      />

      <DataTable
        title="Daily Summary"
        data={snowieOverviewData.dailySummary}
        columns={columns}
        searchable={false}
      />
    </div>
  );
}
