import { useQuery } from "@tanstack/react-query";
import { getCloserxFinancialStats } from "@/services/closerx";

export const useCloserxFinancialStats = (
  startDate?: string,
  endDate?: string | null
) =>
  useQuery({
    queryKey: ["closerx-financials", startDate, endDate],
    queryFn: () =>
      getCloserxFinancialStats({
        startDate,
        endDate,
      }),
  });