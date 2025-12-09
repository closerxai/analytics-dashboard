export const globalOverviewData = {
  kpis: {
    totalUsers: 12847,
    totalRevenue: 458920,
    totalCreditsUsed: 8945320,
    totalInfraCost: 34567
  },
  platformHealth: [
    { name: 'CloserX', status: 'operational' as const },
    { name: 'Snowie', status: 'operational' as const },
    { name: 'Maya', status: 'operational' as const }
  ],
  companyStats: [
    {
      name: 'CloserX',
      users: 5234,
      revenue: 234567,
      calls: 21401,
      credits: 3245890
    },
    {
      name: 'Snowie',
      users: 4512,
      revenue: 189450,
      calls: 23456,
      credits: 3456789
    },
    {
      name: 'Maya',
      users: 3101,
      revenue: 156780,
      calls: 8934,
      credits: 2123450
    }
  ],
  userGrowthTrend: [
    { date: '2024-11-01', value: 11200 },
    { date: '2024-11-05', value: 11450 },
    { date: '2024-11-10', value: 11800 },
    { date: '2024-11-15', value: 12100 },
    { date: '2024-11-20', value: 12400 },
    { date: '2024-11-25', value: 12650 },
    { date: '2024-11-30', value: 12847 }
  ],
  revenueTrend: [
    { date: '2024-11-01', value: 85000 },
    { date: '2024-11-05', value: 125000 },
    { date: '2024-11-10', value: 198000 },
    { date: '2024-11-15', value: 267000 },
    { date: '2024-11-20', value: 345000 },
    { date: '2024-11-25', value: 402000 },
    { date: '2024-11-30', value: 458920 }
  ],
  callMinutesTrend: [
    { date: '2024-11-01', value: 12500 },
    { date: '2024-11-05', value: 15800 },
    { date: '2024-11-10', value: 18900 },
    { date: '2024-11-15', value: 22400 },
    { date: '2024-11-20', value: 26700 },
    { date: '2024-11-25', value: 29800 },
    { date: '2024-11-30', value: 33450 }
  ],
  costTrend: [
    { date: '2024-11-01', value: 5200 },
    { date: '2024-11-05', value: 8900 },
    { date: '2024-11-10', value: 13400 },
    { date: '2024-11-15', value: 18700 },
    { date: '2024-11-20', value: 24200 },
    { date: '2024-11-25', value: 29100 },
    { date: '2024-11-30', value: 34567 }
  ]
};

export const closerxOverviewData = {
  kpis: {
    totalUsers: 5234,
    activeCampaigns: 187,
    quickCalls: 8945,
    campaignCalls: 12456,
    totalMinutes: 15678,
    creditBurn: 3245890
  },
  callsPerDay: [
    { date: '2024-11-24', value: 645 },
    { date: '2024-11-25', value: 712 },
    { date: '2024-11-26', value: 689 },
    { date: '2024-11-27', value: 734 },
    { date: '2024-11-28', value: 798 },
    { date: '2024-11-29', value: 823 },
    { date: '2024-11-30', value: 856 }
  ],
  minutesPerDay: [
    { date: '2024-11-24', value: 1823 },
    { date: '2024-11-25', value: 2145 },
    { date: '2024-11-26', value: 1967 },
    { date: '2024-11-27', value: 2289 },
    { date: '2024-11-28', value: 2456 },
    { date: '2024-11-29', value: 2578 },
    { date: '2024-11-30', value: 2689 }
  ],
  creditBurnTrend: [
    { date: '2024-11-24', value: 98450 },
    { date: '2024-11-25', value: 112340 },
    { date: '2024-11-26', value: 103290 },
    { date: '2024-11-27', value: 118670 },
    { date: '2024-11-28', value: 125890 },
    { date: '2024-11-29', value: 132450 },
    { date: '2024-11-30', value: 138920 }
  ],
  topAgencies: [
    { name: 'Apex Outreach', minutes: 4567 },
    { name: 'Elite Connect', minutes: 3892 },
    { name: 'ProLead Partners', minutes: 3456 },
    { name: 'Summit Sales', minutes: 2987 },
    { name: 'Pinnacle Callers', minutes: 2654 }
  ]
};

export const closerxUsersData = {
  kpis: {
    totalUsers: 5234,
    activeUsers: 3456,
    newUsersYesterday: 23
  },
  userGrowth: [
    { date: '2024-11-01', value: 4750 },
    { date: '2024-11-05', value: 4820 },
    { date: '2024-11-10', value: 4910 },
    { date: '2024-11-15', value: 5020 },
    { date: '2024-11-20', value: 5098 },
    { date: '2024-11-25', value: 5167 },
    { date: '2024-11-30', value: 5234 }
  ],
  users: [
    {
      id: '1',
      name: 'Sarah Johnson',
      quickCalls: 234,
      campaignCalls: 456,
      minutesUsed: 1245,
      creditBurn: 45670,
      lastActive: '2024-11-30T14:23:00Z'
    },
    {
      id: '2',
      name: 'Mike Chen',
      quickCalls: 189,
      campaignCalls: 378,
      minutesUsed: 1089,
      creditBurn: 39850,
      lastActive: '2024-11-30T11:45:00Z'
    },
    {
      id: '3',
      name: 'Emily Rodriguez',
      quickCalls: 312,
      campaignCalls: 524,
      minutesUsed: 1567,
      creditBurn: 57340,
      lastActive: '2024-11-30T16:12:00Z'
    },
    {
      id: '4',
      name: 'David Kim',
      quickCalls: 156,
      campaignCalls: 289,
      minutesUsed: 845,
      creditBurn: 30920,
      lastActive: '2024-11-29T18:34:00Z'
    },
    {
      id: '5',
      name: 'Jessica Martinez',
      quickCalls: 267,
      campaignCalls: 412,
      minutesUsed: 1178,
      creditBurn: 43120,
      lastActive: '2024-11-30T09:56:00Z'
    },
    {
      id: '6',
      name: 'Ryan Thompson',
      quickCalls: 198,
      campaignCalls: 345,
      minutesUsed: 967,
      creditBurn: 35410,
      lastActive: '2024-11-30T13:28:00Z'
    },
    {
      id: '7',
      name: 'Amanda Lee',
      quickCalls: 223,
      campaignCalls: 398,
      minutesUsed: 1123,
      creditBurn: 41090,
      lastActive: '2024-11-30T15:42:00Z'
    },
    {
      id: '8',
      name: 'Chris Anderson',
      quickCalls: 178,
      campaignCalls: 301,
      minutesUsed: 892,
      creditBurn: 32650,
      lastActive: '2024-11-29T21:15:00Z'
    }
  ]
};

export const closerxFinancialsData = {
  kpis: {
    revenue: 234567,
    refunded: 12340,
    disputes_lost: 3450,
    profit: 218777
  },
  paymentsTrend: [
    { date: '2024-11-01', value: 35600 },
    { date: '2024-11-05', value: 52340 },
    { date: '2024-11-10', value: 78920 },
    { date: '2024-11-15', value: 105670 },
    { date: '2024-11-20', value: 145890 },
    { date: '2024-11-25', value: 189340 },
    { date: '2024-11-30', value: 234567 }
  ],
  creditBurnVsPayments: [
    { date: '2024-11-01', creditBurn: 145600, payments: 35600 },
    { date: '2024-11-05', creditBurn: 198340, payments: 52340 },
    { date: '2024-11-10', creditBurn: 267920, payments: 78920 },
    { date: '2024-11-15', creditBurn: 334670, payments: 105670 },
    { date: '2024-11-20', creditBurn: 412890, payments: 145890 },
    { date: '2024-11-25', creditBurn: 478340, payments: 189340 },
    { date: '2024-11-30', creditBurn: 534567, payments: 234567 }
  ],
  revenueByType: [
    { name: 'Quick Campaigns', value: 89450 },
    { name: 'Standard Campaigns', value: 109000 }
  ]
};

export const snowieOverviewData = {
  kpis: {
    totalUsers: 4512,
    totalCalls: 23456,
    totalDuration: 9876,
    totalCredits: 3456789
  },
  callsTrend: [
    { date: '2024-11-24', calls: 2845, minutes: 1234 },
    { date: '2024-11-25', calls: 3012, minutes: 1367 },
    { date: '2024-11-26', calls: 2967, minutes: 1289 },
    { date: '2024-11-27', calls: 3156, minutes: 1445 },
    { date: '2024-11-28', calls: 3289, minutes: 1512 },
    { date: '2024-11-29', calls: 3423, minutes: 1589 },
    { date: '2024-11-30', calls: 3567, minutes: 1623 }
  ],
  modelCredits: [
    { name: 'Vision', value: 1234567 },
    { name: 'Thunder', value: 1456789 },
    { name: 'Avatar', value: 765433 }
  ],
  creditsTrend: [
    { date: '2024-11-24', value: 456789 },
    { date: '2024-11-25', value: 512340 },
    { date: '2024-11-26', value: 489670 },
    { date: '2024-11-27', value: 534890 },
    { date: '2024-11-28', value: 567230 },
    { date: '2024-11-29', value: 598450 },
    { date: '2024-11-30', value: 623120 }
  ],
  dailySummary: [
    { date: '2024-11-24', calls: 2845, duration: 1234, credits: 456789 },
    { date: '2024-11-25', calls: 3012, duration: 1367, credits: 512340 },
    { date: '2024-11-26', calls: 2967, duration: 1289, credits: 489670 },
    { date: '2024-11-27', calls: 3156, duration: 1445, credits: 534890 },
    { date: '2024-11-28', calls: 3289, duration: 1512, credits: 567230 },
    { date: '2024-11-29', calls: 3423, duration: 1589, credits: 598450 },
    { date: '2024-11-30', calls: 3567, duration: 1623, credits: 623120 }
  ]
};

export const snowieUsersData = {
  kpis: {
    totalUsers: 4512,
    activeUsers: 2987,
    newUsersYesterday: 18
  },
  userGrowth: [
    { date: '2024-11-01', value: 4120 },
    { date: '2024-11-05', value: 4178 },
    { date: '2024-11-10', value: 4256 },
    { date: '2024-11-15', value: 4334 },
    { date: '2024-11-20', value: 4401 },
    { date: '2024-11-25', value: 4467 },
    { date: '2024-11-30', value: 4512 }
  ],
  users: [
    {
      id: '1',
      name: 'Alex Turner',
      totalCalls: 456,
      totalDuration: 234,
      creditsUsed: 78900,
      lastActive: '2024-11-30T15:45:00Z'
    },
    {
      id: '2',
      name: 'Sophia Bennett',
      totalCalls: 389,
      totalDuration: 198,
      creditsUsed: 67320,
      lastActive: '2024-11-30T12:23:00Z'
    },
    {
      id: '3',
      name: 'Marcus Johnson',
      totalCalls: 512,
      totalDuration: 267,
      creditsUsed: 89450,
      lastActive: '2024-11-30T17:12:00Z'
    },
    {
      id: '4',
      name: 'Olivia Zhang',
      totalCalls: 298,
      totalDuration: 156,
      creditsUsed: 52340,
      lastActive: '2024-11-29T19:34:00Z'
    },
    {
      id: '5',
      name: 'Ethan Williams',
      totalCalls: 423,
      totalDuration: 212,
      creditsUsed: 71230,
      lastActive: '2024-11-30T10:56:00Z'
    },
    {
      id: '6',
      name: 'Isabella Garcia',
      totalCalls: 367,
      totalDuration: 189,
      creditsUsed: 63410,
      lastActive: '2024-11-30T14:28:00Z'
    },
    {
      id: '7',
      name: 'Noah Davis',
      totalCalls: 445,
      totalDuration: 223,
      creditsUsed: 75090,
      lastActive: '2024-11-30T16:42:00Z'
    },
    {
      id: '8',
      name: 'Emma Martinez',
      totalCalls: 334,
      totalDuration: 178,
      creditsUsed: 59650,
      lastActive: '2024-11-29T22:15:00Z'
    }
  ]
};

export const snowieFinancialsData = {
  kpis: {
    revenue: 189450,
    refunded: 8920,
    disputes_lost: 2340,
    profit: 178190
  },
  paymentsTrend: [
    { date: '2024-11-01', value: 28900 },
    { date: '2024-11-05', value: 45670 },
    { date: '2024-11-10', value: 67340 },
    { date: '2024-11-15', value: 89120 },
    { date: '2024-11-20', value: 123450 },
    { date: '2024-11-25', value: 156780 },
    { date: '2024-11-30', value: 189450 }
  ],
  creditsPurchasedTrend: [
    { date: '2024-11-01', value: 456780 },
    { date: '2024-11-05', value: 689450 },
    { date: '2024-11-10', value: 923400 },
    { date: '2024-11-15', value: 1234560 },
    { date: '2024-11-20', value: 1789230 },
    { date: '2024-11-25', value: 2345670 },
    { date: '2024-11-30', value: 3456789 }
  ],
  revenueByModel: [
    { name: 'Vision', value: 56780 },
    { name: 'Thunder', value: 67450 },
    { name: 'Avatar', value: 43000 }
  ]
};

export const mayaOverviewData = {
  kpis: {
    totalUsers: 3101,
    activeUsers: 1876,
    newUsersYesterday: 12,
    creditsUsed: 2123450,
    callsThisMonth: 8934,
    callDuration: 5678
  },
  userGrowth: [
    { date: '2024-11-01', value: 2850 },
    { date: '2024-11-05', value: 2901 },
    { date: '2024-11-10', value: 2956 },
    { date: '2024-11-15', value: 3003 },
    { date: '2024-11-20', value: 3045 },
    { date: '2024-11-25', value: 3078 },
    { date: '2024-11-30', value: 3101 }
  ],
  creditsTrend: [
    { date: '2024-11-01', value: 298450 },
    { date: '2024-11-05', value: 456780 },
    { date: '2024-11-10', value: 689230 },
    { date: '2024-11-15', value: 923450 },
    { date: '2024-11-20', value: 1234560 },
    { date: '2024-11-25', value: 1678900 },
    { date: '2024-11-30', value: 2123450 }
  ],
  callsTrend: [
    { date: '2024-11-01', value: 1234 },
    { date: '2024-11-05', value: 1567 },
    { date: '2024-11-10', value: 2012 },
    { date: '2024-11-15', value: 2456 },
    { date: '2024-11-20', value: 3123 },
    { date: '2024-11-25', value: 3789 },
    { date: '2024-11-30', value: 4234 }
  ]
};

export const mayaUsersData = {
  kpis: {
    totalUsers: 3101,
    activeUsers: 1876,
    newUsersYesterday: 12
  },
  userGrowth: [
    { date: '2024-11-01', value: 2850 },
    { date: '2024-11-05', value: 2901 },
    { date: '2024-11-10', value: 2956 },
    { date: '2024-11-15', value: 3003 },
    { date: '2024-11-20', value: 3045 },
    { date: '2024-11-25', value: 3078 },
    { date: '2024-11-30', value: 3101 }
  ],
  users: [
    {
      id: '1',
      name: 'James Wilson',
      email: 'james.w@example.com',
      creditsUsed: 45670,
      callsCount: 234,
      lastActive: '2024-11-30T16:23:00Z'
    },
    {
      id: '2',
      name: 'Mia Thompson',
      email: 'mia.t@example.com',
      creditsUsed: 38920,
      callsCount: 198,
      lastActive: '2024-11-30T13:45:00Z'
    },
    {
      id: '3',
      name: 'Lucas Brown',
      email: 'lucas.b@example.com',
      creditsUsed: 52340,
      callsCount: 267,
      lastActive: '2024-11-30T18:12:00Z'
    },
    {
      id: '4',
      name: 'Ava Martinez',
      email: 'ava.m@example.com',
      creditsUsed: 29870,
      callsCount: 156,
      lastActive: '2024-11-29T20:34:00Z'
    },
    {
      id: '5',
      name: 'Oliver Anderson',
      email: 'oliver.a@example.com',
      creditsUsed: 41230,
      callsCount: 212,
      lastActive: '2024-11-30T11:56:00Z'
    },
    {
      id: '6',
      name: 'Charlotte Lee',
      email: 'charlotte.l@example.com',
      creditsUsed: 36710,
      callsCount: 189,
      lastActive: '2024-11-30T15:28:00Z'
    },
    {
      id: '7',
      name: 'William Garcia',
      email: 'william.g@example.com',
      creditsUsed: 44090,
      callsCount: 223,
      lastActive: '2024-11-30T17:42:00Z'
    },
    {
      id: '8',
      name: 'Amelia Davis',
      email: 'amelia.d@example.com',
      creditsUsed: 32650,
      callsCount: 178,
      lastActive: '2024-11-29T23:15:00Z'
    }
  ]
};

export const mayaFinancialsData = {
  kpis: {
    revenue: 156780,
    refunded: 6780,
    disputes_lost: 1890,
    profit: 148110
  },
  paymentsTrend: [
    { date: '2024-11-01', value: 23450 },
    { date: '2024-11-05', value: 36780 },
    { date: '2024-11-10', value: 54230 },
    { date: '2024-11-15', value: 72890 },
    { date: '2024-11-20', value: 98450 },
    { date: '2024-11-25', value: 127890 },
    { date: '2024-11-30', value: 156780 }
  ],
  creditsPurchasedTrend: [
    { date: '2024-11-01', value: 298450 },
    { date: '2024-11-05', value: 456780 },
    { date: '2024-11-10', value: 689230 },
    { date: '2024-11-15', value: 923450 },
    { date: '2024-11-20', value: 1234560 },
    { date: '2024-11-25', value: 1678900 },
    { date: '2024-11-30', value: 2123450 }
  ]
};

export const costsData = {
  kpis: {
    awsCost: 15678,
    gcpCost: 8923,
    twilioCost: 9966,
    totalInfraCost: 34567
  },
  providerBreakdown: [
    { name: 'AWS', value: 15678 },
    { name: 'GCP', value: 8923 },
    { name: 'Twilio/Ultravox', value: 9966 }
  ],
  costTrend: [
    { date: '2024-11-01', value: 5234 },
    { date: '2024-11-05', value: 8967 },
    { date: '2024-11-10', value: 13456 },
    { date: '2024-11-15', value: 18789 },
    { date: '2024-11-20', value: 24123 },
    { date: '2024-11-25', value: 29456 },
    { date: '2024-11-30', value: 34567 }
  ],
  perServiceCosts: [
    { service: 'EC2 Instances', cost: 6789 },
    { service: 'S3 Storage', cost: 2345 },
    { service: 'RDS Database', cost: 4567 },
    { service: 'CloudFront CDN', cost: 1977 },
    { service: 'GCP Compute', cost: 5234 },
    { service: 'GCP Storage', cost: 3689 },
    { service: 'Twilio Voice', cost: 6789 },
    { service: 'Ultravox API', cost: 3177 }
  ],
  costDetails: [
    {
      provider: 'AWS',
      service: 'EC2 Instances',
      spend: 6789,
      change: 12.3,
      notes: 'Increased capacity for CloserX'
    },
    {
      provider: 'AWS',
      service: 'S3 Storage',
      spend: 2345,
      change: 3.2,
      notes: 'Recording storage growth'
    },
    {
      provider: 'AWS',
      service: 'RDS Database',
      spend: 4567,
      change: 8.7,
      notes: 'Database scaling'
    },
    {
      provider: 'AWS',
      service: 'CloudFront CDN',
      spend: 1977,
      change: -2.1,
      notes: 'Optimized caching'
    },
    {
      provider: 'GCP',
      service: 'Compute Engine',
      spend: 5234,
      change: 15.6,
      notes: 'Snowie model inference'
    },
    {
      provider: 'GCP',
      service: 'Cloud Storage',
      spend: 3689,
      change: 5.4,
      notes: 'Model weights storage'
    },
    {
      provider: 'Twilio',
      service: 'Voice Minutes',
      spend: 6789,
      change: 18.9,
      notes: 'Call volume increase'
    },
    {
      provider: 'Ultravox',
      service: 'API Calls',
      spend: 3177,
      change: 11.2,
      notes: 'Maya usage spike'
    }
  ]
};
