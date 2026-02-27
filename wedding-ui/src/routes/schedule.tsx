import { createFileRoute } from '@tanstack/react-router'
import Schedule from '../components/Schedule'

export const Route = createFileRoute('/schedule')({ component: SchedulePage })

function SchedulePage() {
  return (
    <main className="min-h-screen bg-white">
      <Schedule />
    </main>
  )
}
