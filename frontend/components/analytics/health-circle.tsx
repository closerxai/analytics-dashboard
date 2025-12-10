'use client';

import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';

interface HealthCircleProps {
  name: string;
status: "healthy" | "down" | "partial"}

function HealthCircle({
  label,
  status,
  size = "lg",
}: {
  label: string;
  status: "healthy" | "down" | "partial";
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
      ? "w-16 h-16 text-[9px]"
      : "w-20 h-20 text-[10px]";

  return (
    <div
      className={`flex items-center justify-center rounded-full text-white shadow-lg ${dims} ${color}`}
    >
      <span
        className="
          text-center
          leading-tight
          font-medium
          px-1
          line-clamp-2
          break-words
        "
      >
        {label}
      </span>
    </div>
  );
}
