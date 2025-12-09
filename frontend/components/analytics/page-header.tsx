'use client';

import { Download } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { DateRangePicker } from './date-range-picker';
import { DateRange } from 'react-day-picker';

interface PageHeaderProps {
  title: string;
  description?: string;
  showDatePicker?: boolean;
  showExport?: boolean;
  onExport?: () => void;
  onDateChange?: (date: DateRange | undefined) => void;
}

export function PageHeader({
  title,
  description,
  showDatePicker = true,
  showExport = true,
  onExport,
  onDateChange
}: PageHeaderProps) {
  return (
    <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
      <div>
        <h1 className="text-3xl font-semibold tracking-tight text-foreground">
          {title}
        </h1>
        {description && (
          <p className="mt-1 text-sm text-muted-foreground">{description}</p>
        )}
      </div>
      <div className="flex flex-col gap-3 sm:flex-row sm:items-center">
        {showDatePicker && <DateRangePicker onDateChange={onDateChange} />}
        {showExport && (
          <Button variant="outline" onClick={onExport}>
            <Download className="mr-2 h-4 w-4" />
            Export
          </Button>
        )}
      </div>
    </div>
  );
}
