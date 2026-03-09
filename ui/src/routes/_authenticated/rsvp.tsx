import { createFileRoute, useRouteContext } from '@tanstack/react-router'
import RSVP from '../../components/RSVP'

export const Route = createFileRoute('/_authenticated/rsvp')({
  component: RSVPPage,
})

function RSVPPage() {
  const { user } = useRouteContext({ from: '/_authenticated' })
  return (
    <main className="flex grow items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <RSVP user={user} />
    </main>
  )
}
