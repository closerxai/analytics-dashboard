'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { mayaOverviewData } from '@/lib/mock-data';
import { Users, UserCheck, UserPlus, Zap, Phone, Clock } from 'lucide-react';

export default function MayaOverview() {
  return (
    <div className="space-y-8">
      <PageHeader
        title="Maya Overview"
        description="User engagement and credit consumption"
      />

      <KPIGrid columns={3}>
        <MetricCard
          label="Total Users"
          value={mayaOverviewData.kpis.totalUsers}
          icon={Users}
        />
        <MetricCard
          label="Active Users"
          value={mayaOverviewData.kpis.activeUsers}
          icon={UserCheck}
        />
        <MetricCard
          label="New Users Yesterday"
          value={mayaOverviewData.kpis.newUsersYesterday}
          icon={UserPlus}
        />
        <MetricCard
          label="Credits Used"
          value={mayaOverviewData.kpis.creditsUsed}
          icon={Zap}
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Calls This Month"
          value={mayaOverviewData.kpis.callsThisMonth}
          icon={Phone}
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Call Duration"
          value={mayaOverviewData.kpis.callDuration}
          icon={Clock}
          suffix=" min"
          formatValue={(val) => val.toLocaleString()}
        />
      </KPIGrid>

      <div className="grid gap-6 lg:grid-cols-2">
        <LineAreaChart
          title="User Growth"
          data={mayaOverviewData.userGrowth}
          valueFormatter={(val) => val.toLocaleString()}
        />
        <LineAreaChart
          title="Credits Used Trend"
          data={mayaOverviewData.creditsTrend}
          valueFormatter={(val) => val.toLocaleString()}
        />
      </div>

      <LineAreaChart
        title="Calls Trend"
        data={mayaOverviewData.callsTrend}
        valueFormatter={(val) => val.toLocaleString()}
      />
    </div>
  );
}
