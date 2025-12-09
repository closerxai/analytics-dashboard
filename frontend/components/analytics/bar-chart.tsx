'use client';

import {
  Bar,
  BarChart as RechartsBarChart,
  CartesianGrid,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

interface DataPoint {
  [key: string]: string | number;
}

interface BarChartProps {
  title: string;
  data: DataPoint[];
  dataKey: string;
  categoryKey: string;
  color?: string;
  valueFormatter?: (value: number) => string;
}

export function BarChart({
  title,
  data,
  dataKey,
  categoryKey,
  color = 'hsl(var(--primary))',
  valueFormatter = (value) => value.toLocaleString()
}: BarChartProps) {
  return (
    <Card className="border-border/40 bg-card/50 backdrop-blur-sm">
      <CardHeader>
        <CardTitle className="text-base font-medium">{title}</CardTitle>
      </CardHeader>
      <CardContent>
        <ResponsiveContainer width="100%" height={300}>
          <RechartsBarChart data={data}>
            <CartesianGrid
              strokeDasharray="3 3"
              stroke="hsl(var(--border))"
              vertical={false}
            />
            <XAxis
              dataKey={categoryKey}
              stroke="hsl(var(--muted-foreground))"
              fontSize={12}
              tickLine={false}
              axisLine={false}
            />
            <YAxis
              stroke="hsl(var(--muted-foreground))"
              fontSize={12}
              tickLine={false}
              axisLine={false}
              tickFormatter={valueFormatter}
            />
            <Tooltip
              content={({ active, payload }) => {
                if (!active || !payload?.length) return null;
                return (
                  <div className="rounded-lg border border-border bg-background p-3 shadow-lg">
                    <p className="mb-1 text-sm font-medium">
                      {payload[0].payload[categoryKey]}
                    </p>
                    <p className="text-sm text-muted-foreground">
                      {valueFormatter(Number(payload[0].value))}
                    </p>
                  </div>
                );
              }}
            />
            <Bar dataKey={dataKey} fill={color} radius={[4, 4, 0, 0]} />
          </RechartsBarChart>
        </ResponsiveContainer>
      </CardContent>
    </Card>
  );
}
