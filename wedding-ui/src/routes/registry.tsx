import { createFileRoute } from '@tanstack/react-router'
import Registry from '../components/Registry'

export const Route = createFileRoute('/registry')({ component: RegistryPage })

function RegistryPage() {
  return (
    <main className="min-h-screen bg-white">
      <Registry />
    </main>
  )
}
