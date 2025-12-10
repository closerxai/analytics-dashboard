"use client";

import { useState } from "react";
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

function Row({
  label,
  metrics,
  level = 0,
  expandable = false,
  children,
}: {
  label: string;
  metrics: Metrics;
  level?: number;
  expandable?: boolean;
  children?: React.ReactNode;
}) {
  const [open, setOpen] = useState(false);

  return (
    <>
      <tr
        className="cursor-pointer hover:bg-muted"
        onClick={() => expandable && setOpen(!open)}
      >
        <td className="py-2 pl-4" style={{ paddingLeft: level * 20 }}>
          {expandable && (
            <span className="mr-2 text-xs">{open ? "▼" : "▶"}</span>
          )}
          {label}
        </td>

        <td>{formatUSD(metrics.revenue.total)}</td>
        <td>{formatUSD(metrics.refunded.total)}</td>
        <td>{formatUSD(metrics.disputed.total)}</td>
        <td className="font-medium">
          {formatUSD(metrics.profit.total)}
        </td>
      </tr>

      {open && children}
    </>
  );
}

export function FinancialBreakdownTable({ data }: { data: FinancialData }) {
  return (
    <div className="rounded-md border overflow-hidden">
      <table className="w-full text-sm">
        <thead className="bg-muted">
          <tr>
            <th className="text-left p-3">Category</th>
            <th>Revenue</th>
            <th>Refunded</th>
            <th>Disputed</th>
            <th>Profit</th>
          </tr>
        </thead>

        <tbody>
          {/* TOTAL */}
          <Row label="Total" metrics={data.total} expandable>
            {/* Subscriptions */}
            <Row
              label="Subscriptions"
              metrics={data.subscriptions.total}
              level={1}
              expandable
            >
              {Object.entries(data.subscriptions.prices).map(
                ([price, metrics]) => (
                  <Row
                    key={price}
                    label={`$${price}`}
                    metrics={metrics}
                    level={2}
                  />
                )
              )}
            </Row>

            {/* Credits */}
            <Row
              label="Credits"
              metrics={data.credits.total}
              level={1}
              expandable
            >
              {Object.entries(data.credits.prices).map(
                ([price, metrics]) => (
                  <Row
                    key={price}
                    label={`$${price}`}
                    metrics={metrics}
                    level={2}
                  />
                )
              )}
            </Row>

            {/* Others */}
            <Row
              label="Others"
              metrics={data.others.total}
              level={1}
              expandable
            >
              {Object.entries(data.others.prices).map(
                ([price, metrics]) => (
                  <Row
                    key={price}
                    label={`$${price}`}
                    metrics={metrics}
                    level={2}
                  />
                )
              )}
            </Row>
          </Row>
        </tbody>
      </table>
    </div>
  );
}
