"use client";

import { Card } from "@/components/ui/card";

type HealthStatus = "healthy" | "partial" | "down";

type HealthResponse = {
  closerx: boolean;
  snowie: boolean;
  maya: Record<string, boolean>;
};

function statusFromBool(ok: boolean): HealthStatus {
  return ok ? "healthy" : "down";
}

function formatServiceName(key: string) {
  return key
    .replace(/service$/i, "")
    .replace(/([a-z])([A-Z])/g, "$1 $2")
    .replace(/^./, (c) => c.toUpperCase());
}

function HealthCircle({
  label,
  status,
  size = "lg",
}: {
  label: string;
  status: HealthStatus;
  size?: "lg" | "sm" | "xsm";
}) {
  const color =
    status === "healthy"
      ? "bg-green-500/90 shadow-green-500/40"
      : status === "partial"
      ? "bg-yellow-500/90 shadow-yellow-500/40"
      : "bg-red-500/90 shadow-red-500/40";

  const dims =
    size === "lg"
      ? "w-28 h-28 text-lg"
      : size === "xsm"
      ? "w-16 h-16 text-[10px] leading-tight px-1 text-center"
      : "w-20 h-20 text-[10px] leading-tight px-1 text-center";

  return (
    <div
      className={`flex items-center justify-center rounded-full font-semibold text-white shadow-lg ${dims} ${color}`}
    >
      {label}
    </div>
  );
}

export function SystemHealth({ data }: { data: HealthResponse }) {
  const mayaOverall =
    Object.values(data.maya).every(Boolean)
      ? "healthy"
      : Object.values(data.maya).some(Boolean)
      ? "partial"
      : "down";

  return (
    <div>
      <h2 className="mb-4 text-xl font-semibold">System Health</h2>

      <div className="flex items-start gap-12">
        {/* CloserX */}
        <HealthCircle
          label="CloserX"
          status={statusFromBool(data.closerx)}
        />

        {/* Snowie */}
        <HealthCircle
          label="Snowie"
          status={statusFromBool(data.snowie)}
        />

        {/* Maya */}
        <div className="flex items-start gap-6">
          <HealthCircle label="Maya" status={mayaOverall} />

          {/* Maya Services */}
          <Card className="border-border/40 bg-card/50 backdrop-blur-sm p-4">
            <div className="grid grid-cols-4 gap-3">
              {Object.entries(data.maya).map(([service, ok]) => {
                const name = formatServiceName(service);

                return (
                  <HealthCircle
                    key={service}
                    label={name}
                    status={ok ? "healthy" : "down"}
                    size="xsm"
                  />
                );
              })}
            </div>
          </Card>
        </div>
      </div>
    </div>
  );
}
