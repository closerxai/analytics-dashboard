import './globals.css';
import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import { MainLayout } from '@/components/layout/main-layout';
import { QueryProvider } from '@/providers/QueryProvider';
import { Toaster } from "@/components/ui/sonner";


const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Analytics Dashboard',
  description: 'Internal analytics dashboard for CloserX, Snowie, and Maya',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="dark">
      <body className={inter.className}>
        <QueryProvider>
          <MainLayout>{children}</MainLayout>
          <Toaster />
        </QueryProvider>
      </body>
    </html>
  );
}
