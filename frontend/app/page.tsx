'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { HealthCircle } from '@/components/analytics/health-circle';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { globalOverviewData } from '@/lib/mock-data';
import { Users, DollarSign, Zap, Server, Phone } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

export default function GlobalOverview() {
  return (
    <div className="space-y-8">
      <PageHeader
        title="Global Overview"
        description="Cross-platform analytics and system health"
      />

      <KPIGrid columns={4}>
        <MetricCard
          label="Total Users"
          value={globalOverviewData.kpis.totalUsers}
          icon={Users}
        />
        <MetricCard
          label="Total Revenue"
          value={globalOverviewData.kpis.totalRevenue}
          icon={DollarSign}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Total Credits Used"
          value={globalOverviewData.kpis.totalCreditsUsed}
          icon={Zap}
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Total Infra Cost"
          value={globalOverviewData.kpis.totalInfraCost}
          icon={Server}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
      </KPIGrid>

      <div>
        <h2 className="mb-4 text-xl font-semibold">System Health</h2>
        <div className="flex flex-wrap gap-8">
          {globalOverviewData.platformHealth.map((platform) => (
            <HealthCircle
              key={platform.name}
              name={platform.name}
              status={platform.status}
            />
          ))}
        </div>
      </div>

      <div>
        <h2 className="mb-4 text-xl font-semibold">Company Quick Stats</h2>
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {globalOverviewData.companyStats.map((company) => (
            <Card
              key={company.name}
              className="border-border/40 bg-card/50 backdrop-blur-sm"
            >
              <CardHeader>
                <CardTitle className="text-lg">{company.name}</CardTitle>
              </CardHeader>
              <CardContent className="space-y-3">
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <Users className="h-4 w-4" />
                    <span>Users</span>
                  </div>
                  <span className="text-lg font-semibold">
                    {company.users.toLocaleString()}
                  </span>
                </div>
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <DollarSign className="h-4 w-4" />
                    <span>Revenue</span>
                  </div>
                  <span className="text-lg font-semibold">
                    ${company.revenue.toLocaleString()}
                  </span>
                </div>
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <Phone className="h-4 w-4" />
                    <span>Calls</span>
                  </div>
                  <span className="text-lg font-semibold">
                    {company.calls.toLocaleString()}
                  </span>
                </div>
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <Zap className="h-4 w-4" />
                    <span>Credits</span>
                  </div>
                  <span className="text-lg font-semibold">
                    {company.credits.toLocaleString()}
                  </span>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>

      <div className="grid gap-6 lg:grid-cols-2">
        <LineAreaChart
          title="User Growth Trend"
          data={globalOverviewData.userGrowthTrend}
          valueFormatter={(val) => val.toLocaleString()}
        />
        <LineAreaChart
          title="Revenue Trend"
          data={globalOverviewData.revenueTrend}
          valueFormatter={(val) => `$${val.toLocaleString()}`}
        />
      </div>

      <div className="grid gap-6 lg:grid-cols-2">
        <LineAreaChart
          title="Call Minutes Trend"
          data={globalOverviewData.callMinutesTrend}
          valueFormatter={(val) => `${val.toLocaleString()} min`}
        />
        <LineAreaChart
          title="Infrastructure Cost Trend"
          data={globalOverviewData.costTrend}
          valueFormatter={(val) => `$${val.toLocaleString()}`}
        />
      </div>
    </div>
  );
}
