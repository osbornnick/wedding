import { createFileRoute } from '@tanstack/react-router'
import Registry from '../../components/Registry'

export const Route = createFileRoute('/_authenticated/registry')({
  component: RegistryPage,
})

function RegistryPage() {
  return (
    <main className="flex grow">
      <Registry />
    </main>
  )
}
