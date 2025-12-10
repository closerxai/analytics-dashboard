import { useQuery } from "@tanstack/react-query";
import { getAWSBilling, getHealthCheckStatus } from "@/services/dashboard";

export const useHealthCheck = () =>
  useQuery({
    queryKey: ["health-check"],
    queryFn: () =>
      getHealthCheckStatus(),
  });
export const useAWSBilling = () =>
  useQuery({
    queryKey: ["aws-billing"],
    queryFn: () =>
      getAWSBilling(),
  });