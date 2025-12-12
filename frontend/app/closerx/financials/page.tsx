"use client";

import { PageHeader } from "@/components/analytics/page-header";
import { LineAreaChart } from "@/components/analytics/line-area-chart";
import { closerxFinancialsData } from "@/lib/mock-data";
import { useCloserxFinancialStats } from "@/hooks/closerx";
import { useEffect, useState } from "react";
import { toast } from "sonner";
import { Skeleton } from "@/components/ui/skeleton";
import { FinancialBreakdownCards } from "@/components/analytics/financial-breakdown-table";
import { DateRange } from "react-day-picker";
import { format } from "date-fns";

export default function CloserXFinancials() {
  const [dateRange, setDateRange] = useState<DateRange | undefined>({
    from: new Date(new Date().getFullYear(), new Date().getMonth(), 1),
    to: new Date(),
  });

  const startDate = dateRange?.from
    ? format(dateRange.from, "yyyy-MM-dd")
    : undefined;
  const endDate = dateRange?.to
    ? format(dateRange.to, "yyyy-MM-dd")
    : undefined;

  const { data, isLoading, isError } = useCloserxFinancialStats(
    startDate,
    endDate
  );

  useEffect(() => {
    if (isError) {
      toast.error("Failed to load financial stats");
    }
  }, [isError]);

  const handleDateChange = (date: DateRange | undefined) => {
    setDateRange(date);
  };

  return (
    <div className="space-y-8">
      <PageHeader
        title="CloserX Financials"
        description="Revenue, subscriptions, credits, and dispute analytics"
        onDateChange={handleDateChange}
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
