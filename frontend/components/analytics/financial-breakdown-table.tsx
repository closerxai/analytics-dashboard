"use client";

import { useState } from "react";
import { ChevronDown, ChevronRight } from "lucide-react";
import { formatUSD } from "@/lib/utils";

type Metric = {
  count: number;
  total: number;
};

type Metrics = {
  revenue: Metric;
  refunded: Metric;
  disputed: Metric;
  profit: Metric;
};

type Section = {
  total: Metrics;
  prices: Record<string, Metrics>;
};

type FinancialData = {
  total: Metrics;
  subscriptions: Section;
  credits: Section;
  others: Section;
};

function MetricCard({
  title,
  metric,
}: {
  title: string;
  metric: Metric;
}) {
  return (
    <div className="flex-1 rounded-lg bg-white/10 backdrop-blur-md p-4">
      <div className="text-xs text-muted-foreground">{title}</div>
      <div className="text-lg font-semibold">{formatUSD(metric.total)}</div>
      <div className="text-xs text-muted-foreground">
        {metric.count} tx
      </div>
    </div>
  );
}

function MetricRow({
  label,
  metrics,
  level = 0,
  children,
}: {
  label: string;
  metrics: Metrics;
  level?: number;
  children?: React.ReactNode;
}) {
  const [open, setOpen] = useState(false);

  return (
    <div className={level > 0 ? "ml-6 mt-3" : "mt-4"}>
      {/* ROW */}
      <div
        onClick={() => setOpen(!open)}
        className="cursor-pointer rounded-xl border bg-white/5 hover:bg-white/10 transition p-4"
      >
        <div className="flex items-center justify-between mb-3">
          <div className="flex items-center gap-2 font-medium">
            {open ? <ChevronDown size={16} /> : <ChevronRight size={16} />}
            {label}
          </div>
        </div>

        <div className="flex gap-3">
          <MetricCard title="Revenue" metric={metrics.revenue} />
          <MetricCard title="Disputed" metric={metrics.disputed} />
          <MetricCard title="Refunded" metric={metrics.refunded} />
          <MetricCard title="Profit" metric={metrics.profit} />
        </div>
      </div>

      {/* CHILDREN */}
      {open && <div>{children}</div>}
    </div>
  );
}

export function FinancialBreakdownCards({
  data,
}: {
  data: FinancialData;
}) {
  return (
    <div>
      <MetricRow label="Total" metrics={data.total}>
        <MetricRow
          label="Subscriptions"
          metrics={data.subscriptions.total}
          level={1}
        >
          {Object.entries(data.subscriptions.prices).map(
            ([price, metrics]) => (
              <MetricRow
                key={price}
                label={`$${price}`}
                metrics={metrics}
                level={2}
              />
            )
          )}
        </MetricRow>

        <MetricRow
          label="Credits"
          metrics={data.credits.total}
          level={1}
        >
          {Object.entries(data.credits.prices).map(
            ([price, metrics]) => (
              <MetricRow
                key={price}
                label={`$${price}`}
                metrics={metrics}
                level={2}
              />
            )
          )}
        </MetricRow>

        <MetricRow
          label="Others"
          metrics={data.others.total}
          level={1}
        >
          {Object.entries(data.others.prices).map(
            ([price, metrics]) => (
              <MetricRow
                key={price}
                label={`$${price}`}
                metrics={metrics}
                level={2}
              />
            )
          )}
        </MetricRow>
      </MetricRow>
    </div>
  );
}
