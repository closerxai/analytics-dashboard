'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { DataTable, Column } from '@/components/analytics/data-table';
import { snowieUsersData } from '@/lib/mock-data';
import { Users, UserCheck, UserPlus } from 'lucide-react';
import { Badge } from '@/components/ui/badge';

export default function SnowieUsers() {
  const columns: Column<typeof snowieUsersData.users[0]>[] = [
    { key: 'name', label: 'Name', sortable: true },
    {
      key: 'totalCalls',
      label: 'Total Calls',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
    {
      key: 'totalDuration',
      label: 'Duration (min)',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
    {
      key: 'creditsUsed',
      label: 'Credits Used',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
    {
      key: 'lastActive',
      label: 'Last Active',
      sortable: true,
      render: (val) => {
        const date = new Date(val);
        const now = new Date();
        const diffHours = Math.floor((now.getTime() - date.getTime()) / (1000 * 60 * 60));

        if (diffHours < 24) {
          return <Badge variant="secondary">Active today</Badge>;
        } else {
          return date.toLocaleDateString();
        }
      },
    },
  ];

  return (
    <div className="space-y-8">
      <PageHeader
        title="Snowie Users"
        description="User activity and model usage patterns"
      />

      <KPIGrid columns={3}>
        <MetricCard
          label="Total Users"
          value={snowieUsersData.kpis.totalUsers}
          icon={Users}
        />
        <MetricCard
          label="Active Users"
          value={snowieUsersData.kpis.activeUsers}
          icon={UserCheck}
        />
        <MetricCard
          label="New Users Yesterday"
          value={snowieUsersData.kpis.newUsersYesterday}
          icon={UserPlus}
        />
      </KPIGrid>

      <LineAreaChart
        title="User Growth"
        data={snowieUsersData.userGrowth}
        valueFormatter={(val) => val.toLocaleString()}
      />

      <DataTable
        title="User Activity"
        data={snowieUsersData.users}
        columns={columns}
        searchPlaceholder="Search by name..."
      />
    </div>
  );
}
