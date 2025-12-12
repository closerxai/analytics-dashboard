"use client";
import {ReactNode, useState} from "react";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";

export function QueryProvider({children}: {children: ReactNode}) {
    const [queryClient] = useState(() => new QueryClient({
        defaultOptions: {
            queries: {
                refetchOnWindowFocus: true,
                retry: false,
            },
        },
    }));
    return (
        <QueryClientProvider client={queryClient}>
            {children}
        </QueryClientProvider>
    );
}