import { api } from "@/lib/axios";

export const getSnowieFinancialStats = async (
    params: {
        startDate?: string;
        endDate?: string | null;
    }
) => {
    const response = await api.get("/snowie/financials", { params });
    return response.data.data;
}