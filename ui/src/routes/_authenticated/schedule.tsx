import { createFileRoute } from '@tanstack/react-router'
import Schedule from '../../components/Schedule'

export const Route = createFileRoute('/_authenticated/schedule')({
  component: SchedulePage,
})

function SchedulePage() {
  return (
    <main className="flex grow">
      <Schedule />
    </main>
  )
}
