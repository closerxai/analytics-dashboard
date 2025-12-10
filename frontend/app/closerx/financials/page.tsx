"use client";

import { PageHeader } from "@/components/analytics/page-header";
import { MetricCard } from "@/components/analytics/metric-card";
import { KPIGrid } from "@/components/analytics/kpi-grid";
import { LineAreaChart } from "@/components/analytics/line-area-chart";
import { closerxFinancialsData } from "@/lib/mock-data";
import { DollarSign, TrendingUp, AlertCircle } from "lucide-react";
import { useCloserxFinancialStats } from "@/hooks/closerx";
import { useEffect } from "react";
import { toast } from "sonner";
import { Skeleton } from "@/components/ui/skeleton";
import { formatUSD } from "@/lib/utils";
import { FinancialBreakdownTable } from "@/components/analytics/financial-breakdown-table";

export default function CloserXFinancials() {
  const { data, isLoading, isError } = useCloserxFinancialStats();

  // Toast on error
  useEffect(() => {
    if (isError) {
      toast.error("Failed to load financial stats", {
        description: "Showing fallback values instead.",
      });
    }
  }, [isError]);

  // ✅ New API shape
  const totals = data?.total;

  const stats = {
    revenue: isError ? "-" : totals?.revenue.total ?? 0,
    refunded: isError ? "-" : totals?.refunded.total ?? 0,
    disputes_lost: isError ? "-" : totals?.disputed.total ?? 0,
    profit: isError ? "-" : totals?.profit.total ?? 0,
  };

  return (
    <div className="space-y-8">
      <PageHeader
        title="CloserX Financials"
        description="Revenue, payments, and credit analytics"
      />

      {/* KPI section */}
      {isLoading ? (
        <KPIGrid columns={4}>
          <Skeleton className="h-[120px] w-full rounded-md" />
          <Skeleton className="h-[120px] w-full rounded-md" />
          <Skeleton className="h-[120px] w-full rounded-md" />
          <Skeleton className="h-[120px] w-full rounded-md" />
        </KPIGrid>
      ) : (
        <KPIGrid columns={4}>
          <MetricCard
            label="Revenue"
            value={stats.revenue}
            icon={DollarSign}
            formatValue={() => (isError ? "-" : formatUSD(stats.revenue))}
          />

          <MetricCard
            label="Refunded"
            value={stats.refunded}
            icon={AlertCircle}
            formatValue={() => (isError ? "-" : formatUSD(stats.refunded))}
          />

          <MetricCard
            label="Disputes Lost"
            value={stats.disputes_lost}
            icon={AlertCircle}
            formatValue={() =>
              isError ? "-" : formatUSD(stats.disputes_lost)
            }
          />

          <MetricCard
            label="Profit"
            value={stats.profit}
            icon={TrendingUp}
            formatValue={() => (isError ? "-" : formatUSD(stats.profit))}
          />
        </KPIGrid>
      )}

      {/* Chart (mock data for now) */}
      <LineAreaChart
        title="Payments Trend"
        data={closerxFinancialsData.paymentsTrend}
        valueFormatter={(val) => `$${val.toLocaleString()}`}
      />

      {/* ✅ Expandable financial breakdown */}
      {!isLoading && !isError && data?.data && (
        <FinancialBreakdownTable data={data.data} />
      )}
    </div>
  );
}
