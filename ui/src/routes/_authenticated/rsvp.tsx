import { createFileRoute } from '@tanstack/react-router'
import RSVP from '../../components/RSVP'

export const Route = createFileRoute('/_authenticated/rsvp')({ component: RSVPPage })

function RSVPPage() {
  return (
    <main className="min-h-screen bg-white">
      <RSVP />
    </main>
  )
}
