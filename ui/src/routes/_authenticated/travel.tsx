import { createFileRoute } from '@tanstack/react-router'
import Travel from '../../components/Travel'

export const Route = createFileRoute('/_authenticated/travel')({
  component: TravelPage,
})

function TravelPage() {
  return (
    <main className="flex grow">
      <Travel />
    </main>
  )
}
