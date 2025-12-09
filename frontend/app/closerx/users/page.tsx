'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { DataTable, Column } from '@/components/analytics/data-table';
import { closerxUsersData } from '@/lib/mock-data';
import { Users, UserCheck, UserPlus } from 'lucide-react';
import { Badge } from '@/components/ui/badge';

export default function CloserXUsers() {
  const columns: Column<typeof closerxUsersData.users[0]>[] = [
    { key: 'name', label: 'Name', sortable: true },
    {
      key: 'quickCalls',
      label: 'Quick Calls',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
    {
      key: 'campaignCalls',
      label: 'Campaign Calls',
      sortable: true,
      render: (val) => val.toLocaleString(),
    },
    {
      key: 'minutesUsed',
      label: 'Minutes Used',
      sortable: true,
      render: (val) => `${val.toLocaleString()} min`,
    },
    {
      key: 'creditBurn',
      label: 'Credit Burn',
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
        title="CloserX Users"
        description="User activity and engagement metrics"
      />

      <KPIGrid columns={3}>
        <MetricCard
          label="Total Users"
          value={closerxUsersData.kpis.totalUsers}
          icon={Users}
        />
        <MetricCard
          label="Active Users"
          value={closerxUsersData.kpis.activeUsers}
          icon={UserCheck}
        />
        <MetricCard
          label="New Users Yesterday"
          value={closerxUsersData.kpis.newUsersYesterday}
          icon={UserPlus}
        />
      </KPIGrid>

      <LineAreaChart
        title="User Growth"
        data={closerxUsersData.userGrowth}
        valueFormatter={(val) => val.toLocaleString()}
      />

      <DataTable
        title="User Activity"
        data={closerxUsersData.users}
        columns={columns}
        searchPlaceholder="Search by name..."
      />
    </div>
  );
}
