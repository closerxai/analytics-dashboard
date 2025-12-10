import { useQuery } from "@tanstack/react-query";
import { getSnowieFinancialStats } from "@/services/snowie";

export const useSnowieFinancialStats = (
  startDate?: string,
  endDate?: string | null
) =>
  useQuery({
    queryKey: ["snowie-financials", startDate, endDate],
    queryFn: () =>
      getSnowieFinancialStats({
        startDate,
        endDate,
      }),
  });