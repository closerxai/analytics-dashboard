'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { DataTable, Column } from '@/components/analytics/data-table';
import { mayaUsersData } from '@/lib/mock-data';
import { Users, UserCheck, UserPlus } from 'lucide-react';
import { Badge } from '@/components/ui/badge';

export default function MayaUsers() {
  const columns: Column<typeof mayaUsersData.users[0]>[] = [
    { key: 'name', label: 'Name', sortable: true },
    { key: 'email', label: 'Email', sortable: true },
    {
      key: 'creditsUsed',
      label: 'Credits Used',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
    {
      key: 'callsCount',
      label: 'Calls',
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
        title="Maya Users"
        description="User activity and engagement metrics"
      />

      <KPIGrid columns={3}>
        <MetricCard
          label="Total Users"
          value={mayaUsersData.kpis.totalUsers}
          icon={Users}
        />
        <MetricCard
          label="Active Users"
          value={mayaUsersData.kpis.activeUsers}
          icon={UserCheck}
        />
        <MetricCard
          label="New Users Yesterday"
          value={mayaUsersData.kpis.newUsersYesterday}
          icon={UserPlus}
        />
      </KPIGrid>

      <LineAreaChart
        title="User Growth"
        data={mayaUsersData.userGrowth}
        valueFormatter={(val) => val.toLocaleString()}
      />

      <DataTable
        title="User Activity"
        data={mayaUsersData.users}
        columns={columns}
        searchPlaceholder="Search by name or email..."
      />
    </div>
  );
}
