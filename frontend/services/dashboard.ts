import { api } from "@/lib/axios";

export const getHealthCheckStatus = async () => {
    const response = await api.get("/health");
    return response.data.data;
}
export const getAWSBilling = async () => {
    const response = await api.get("/aws_billing");
    return response.data.data;
}