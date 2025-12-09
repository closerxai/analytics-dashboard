'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';
import {
  LayoutDashboard,
  Users,
  DollarSign,
  ChevronDown,
  Server,
} from 'lucide-react';
import { useState } from 'react';

interface NavItem {
  title: string;
  href: string;
  icon: React.ComponentType<{ className?: string }>;
}

interface NavSection {
  title: string;
  items: NavItem[];
}

const navSections: NavSection[] = [
  {
    title: 'Overview',
    items: [
      {
        title: 'Global Overview',
        href: '/',
        icon: LayoutDashboard,
      },
    ],
  },
  {
    title: 'CloserX',
    items: [
      {
        title: 'Overview',
        href: '/closerx/overview',
        icon: LayoutDashboard,
      },
      {
        title: 'Users',
        href: '/closerx/users',
        icon: Users,
      },
      {
        title: 'Financials',
        href: '/closerx/financials',
        icon: DollarSign,
      },
    ],
  },
  {
    title: 'Snowie',
    items: [
      {
        title: 'Overview',
        href: '/snowie/overview',
        icon: LayoutDashboard,
      },
      {
        title: 'Users',
        href: '/snowie/users',
        icon: Users,
      },
      {
        title: 'Financials',
        href: '/snowie/financials',
        icon: DollarSign,
      },
    ],
  },
  {
    title: 'Maya',
    items: [
      {
        title: 'Overview',
        href: '/maya/overview',
        icon: LayoutDashboard,
      },
      {
        title: 'Users',
        href: '/maya/users',
        icon: Users,
      },
      {
        title: 'Financials',
        href: '/maya/financials',
        icon: DollarSign,
      },
    ],
  },
  {
    title: 'Infrastructure',
    items: [
      {
        title: 'Costs',
        href: '/costs',
        icon: Server,
      },
    ],
  },
];

function NavSection({ section }: { section: NavSection }) {
  const pathname = usePathname();
  const [isOpen, setIsOpen] = useState(true);

  const hasActiveItem = section.items.some((item) => pathname === item.href);

  return (
    <div className="mb-4">
      <Button
        variant="ghost"
        onClick={() => setIsOpen(!isOpen)}
        className="mb-1 w-full justify-between px-3 py-2 text-xs font-semibold uppercase tracking-wider text-muted-foreground hover:text-foreground"
      >
        {section.title}
        <ChevronDown
          className={cn(
            'h-4 w-4 transition-transform',
            isOpen && 'rotate-180'
          )}
        />
      </Button>
      {isOpen && (
        <div className="space-y-1">
          {section.items.map((item) => {
            const isActive = pathname === item.href;
            const Icon = item.icon;

            return (
              <Link key={item.href} href={item.href}>
                <Button
                  variant={isActive ? 'secondary' : 'ghost'}
                  className={cn(
                    'w-full justify-start gap-3 px-3',
                    isActive && 'bg-secondary font-medium'
                  )}
                >
                  <Icon className="h-4 w-4" />
                  {item.title}
                </Button>
              </Link>
            );
          })}
        </div>
      )}
    </div>
  );
}

export function SidebarNav() {
  return (
    <div className="flex h-full flex-col border-r border-border/40 bg-card/30 backdrop-blur-sm">
      <div className="p-6">
        <h2 className="text-lg font-semibold tracking-tight">Analytics</h2>
        <p className="text-sm text-muted-foreground">Internal Dashboard</p>
      </div>
      <ScrollArea className="flex-1 px-3">
        {navSections.map((section) => (
          <NavSection key={section.title} section={section} />
        ))}
      </ScrollArea>
    </div>
  );
}
