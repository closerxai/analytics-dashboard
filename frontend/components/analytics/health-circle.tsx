'use client';

import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';

interface HealthCircleProps {
  name: string;
  status: 'operational' | 'down';
}

export function HealthCircle({ name, status }: HealthCircleProps) {
  const statusColors = {
    operational: 'bg-green-500 shadow-green-500/50',
    down: 'bg-red-500 shadow-red-500/50'
  };

  const statusMessages = {
    operational: 'All systems operational',
    down: 'Service down'
  };

  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <div
            className={`flex h-32 w-32 cursor-pointer items-center justify-center rounded-full border-4 border-background ${statusColors[status]} shadow-lg transition-all hover:scale-105`}
          >
            <span className="text-base font-semibold text-white">{name}</span>
          </div>
        </TooltipTrigger>
        <TooltipContent>
          <p className="font-medium">{name}</p>
          <p className="text-sm text-muted-foreground">{statusMessages[status]}</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
