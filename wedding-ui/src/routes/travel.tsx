import { createFileRoute } from '@tanstack/react-router'
import Travel from '../components/Travel'

export const Route = createFileRoute('/travel')({ component: TravelPage })

function TravelPage() {
  return (
    <main className="min-h-screen bg-white">
      <Travel />
    </main>
  )
}
