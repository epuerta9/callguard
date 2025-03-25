# CallGuard - Voice Agent Protection

CallGuard is a Next.js application that helps users create voice agents to act as executive assistants for their loved ones. The platform offers protection against scams and malicious calls through customizable guardrails.

## Features

- Create and manage voice agents for your relatives
- Assign phone numbers to agents
- Define custom guardrails for protection
- Track call history and transcripts
- Receive alerts for potentially malicious calls

## Tech Stack

- **Framework**: Next.js with App Router
- **UI**: Tailwind CSS with Shadcn UI components
- **Database**: Supabase
- **Authentication**: Supabase Auth
- **Styling**: Tailwind CSS
- **Deployment**: Vercel

## Getting Started

### Prerequisites

- Node.js 16+
- npm or yarn

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/callguard.git
   cd callguard
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Set up environment variables:
   Create a `.env.local` file with the following:
   ```
   NEXT_PUBLIC_SUPABASE_URL=your_supabase_url
   NEXT_PUBLIC_SUPABASE_ANON_KEY=your_supabase_anon_key
   ```

4. Run the development server:
   ```bash
   npm run dev
   ```

5. Open [http://localhost:3000](http://localhost:3000) to view the application.

## Project Structure

- `/src/app` - Next.js App Router pages
- `/src/components` - Reusable UI components
- `/src/utils` - Utility functions and Supabase clients
- `/src/services` - API service functions
- `/src/types` - TypeScript type definitions

## Database Schema

### Agents
- id (uuid)
- relativeName (string)
- phoneNumber (string)
- guardrails (string array)
- createdAt (timestamp)
- updatedAt (timestamp)

### CallHistory
- id (uuid)
- agentId (uuid, foreign key to Agents)
- callTime (timestamp)
- duration (number)
- transcript (text)
- isMalicious (boolean)
- callerId (string)
- callerName (string, optional)

### MaliciousCalls
- id (uuid)
- callHistoryId (uuid, foreign key to CallHistory)
- reason (string)
- detectedAt (timestamp)
- isResolved (boolean)

## License

This project is licensed under the MIT License - see the LICENSE file for details. 