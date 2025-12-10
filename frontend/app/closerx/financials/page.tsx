"use client";

import { PageHeader } from "@/components/analytics/page-header";
import { LineAreaChart } from "@/components/analytics/line-area-chart";
import { closerxFinancialsData } from "@/lib/mock-data";
import { useCloserxFinancialStats } from "@/hooks/closerx";
import { useEffect } from "react";
import { toast } from "sonner";
import { Skeleton } from "@/components/ui/skeleton";
import { FinancialBreakdownCards } from "@/components/analytics/financial-breakdown-table";

export default function CloserXFinancials() {
  const { data, isLoading, isError } = useCloserxFinancialStats();

  useEffect(() => {
    if (isError) {
      toast.error("Failed to load financial stats");
    }
  }, [isError]);

  return (
    <div className="space-y-8">
      <PageHeader
        title="CloserX Financials"
        description="Revenue, subscriptions, credits, and dispute analytics"
      />

      {/* Loading */}
      {isLoading && (
        <Skeleton className="h-[240px] w-full rounded-md" />
      )}

      {/* ✅ Financial Breakdown (Totals + Expandable Rows) */}
      {!isLoading && !isError && data && (
        <FinancialBreakdownCards data={data} />
      )}

      {/* Chart (optional / mock for now) */}
      <LineAreaChart
        title="Payments Trend"
        data={closerxFinancialsData.paymentsTrend}
        valueFormatter={(val) => `$${val.toLocaleString()}`}
      />
    </div>
  );
}
