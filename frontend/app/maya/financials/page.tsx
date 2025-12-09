'use client';

import { PageHeader } from '@/components/analytics/page-header';
import { MetricCard } from '@/components/analytics/metric-card';
import { KPIGrid } from '@/components/analytics/kpi-grid';
import { LineAreaChart } from '@/components/analytics/line-area-chart';
import { mayaFinancialsData } from '@/lib/mock-data';
import { DollarSign, TrendingUp, AlertCircle, CreditCard } from 'lucide-react';

export default function MayaFinancials() {
  return (
    <div className="space-y-8">
      <PageHeader
        title="Maya Financials"
        description="Revenue, payments, and credit analytics"
      />

      <KPIGrid columns={4}>
        <MetricCard
          label="Revenue"
          value={mayaFinancialsData.kpis.revenue}
          icon={DollarSign}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Refunded"
          value={mayaFinancialsData.kpis.refunded}
          icon={AlertCircle}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Disputes Lost"
          value={mayaFinancialsData.kpis.disputes_lost}
          icon={AlertCircle}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
        <MetricCard
          label="Profit"
          value={mayaFinancialsData.kpis.profit}
          icon={TrendingUp}
          prefix="$"
          formatValue={(val) => val.toLocaleString()}
        />
      </KPIGrid>

      <div className="grid gap-6 lg:grid-cols-2">
        <LineAreaChart
          title="Payments Trend"
          data={mayaFinancialsData.paymentsTrend}
          valueFormatter={(val) => `$${val.toLocaleString()}`}
        />
        <LineAreaChart
          title="Credits Purchased Trend"
          data={mayaFinancialsData.creditsPurchasedTrend}
          valueFormatter={(val) => val.toLocaleString()}
        />
      </div>
    </div>
  );
}
