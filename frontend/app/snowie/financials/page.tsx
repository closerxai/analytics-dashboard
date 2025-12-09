'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { PieChart } from '@/components/analytics/pie-chart';
import { snowieFinancialsData } from '@/lib/mock-data';
import { DollarSign, TrendingUp, AlertCircle, CreditCard } from 'lucide-react';

export default function SnowieFinancials() {
  return (
    <div className="space-y-8">
      <PageHeader
        title="Snowie Financials"
        description="Revenue, payments, and credit analytics"
      />

      <KPIGrid columns={4}>
        <MetricCard
          label="Revenue"
          value={snowieFinancialsData.kpis.revenue}
          icon={DollarSign}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Refunded"
          value={snowieFinancialsData.kpis.refunded}
          icon={AlertCircle}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Disputes Lost"
          value={snowieFinancialsData.kpis.disputes_lost}
          icon={AlertCircle}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Profit"
          value={snowieFinancialsData.kpis.profit}
          icon={TrendingUp}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
      </KPIGrid>

      <div className="grid gap-6 lg:grid-cols-2">
        <LineAreaChart
          title="Payments Trend"
          data={snowieFinancialsData.paymentsTrend}
          valueFormatter={(val) => `$${val.toLocaleString()}`}
        />
        <LineAreaChart
          title="Credits Purchased Trend"
          data={snowieFinancialsData.creditsPurchasedTrend}
          valueFormatter={(val) => val.toLocaleString()}
        />
      </div>

      <PieChart
        title="Revenue by Model"
        data={snowieFinancialsData.revenueByModel}
        valueFormatter={(val) => `$${val.toLocaleString()}`}
      />
    </div>
  );
}
