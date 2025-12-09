import { api } from "@/lib/axios";

export const getCloserxFinancialStats = async (
    params: {
        startDate?: string;
        endDate?: string | null;
    }
) => {
    const response = await api.get("/closerx/financials", { params });
    return response.data.data;
}