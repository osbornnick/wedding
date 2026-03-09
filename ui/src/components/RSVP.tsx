import type { User } from '#/lib/session'
import { useState } from 'react'
import {
  Button,
  Card,
  HelperText,
  Label,
  Spinner,
  TextInput,
} from 'flowbite-react'
import { findGuestsFn } from '#/lib/guests'
import type { Guest } from '#/lib/guests'

export default function RSVP({ user }: { user?: User | null | undefined }) {
  const [submitted, setSubmitted] = useState(user?.hasRSVPed ?? false)
  const [possibleGuests, setPossibleGuests] = useState<Guest[]>([])
  const [name, setName] = useState('')
  const [loading, setLoading] = useState(false)

  const findGuests = async (n: string) => {
    setLoading(true)
    if (!n || n.trim() === '' || n.trim().length < 2) {
      setLoading(false)
      return
    }
    const guests = await findGuestsFn({ data: { name: n } })
    if (guests) {
      setPossibleGuests(guests)
    } else {
      setPossibleGuests([])
    }
    setLoading(false)
  }
  return (
    <section className="py-12 px-4 sm:px-6 lg:px-8 max-w-xl mx-auto">
      <h2 className="mb-4 text-center font-serif text-3xl font-semibold">
        RSVP
      </h2>
      <p className="mb-8 text-center">
        Please respond by <span className="font-medium">April 1, 2027</span>.
      </p>

      {submitted ? (
        <div className="rounded-2xl border border-green-200 bg-green-50 p-8 text-center">
          <p className="text-lg font-medium text-green-800">
            Thank you for your RSVP! We can't wait to celebrate with you.
          </p>
        </div>
      ) : (
        <form>
          <Card className="">
            <Label htmlFor="name">Name</Label>
            <TextInput
              placeholder="Sarah King"
              id="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
            <HelperText className="font-medium max-w-md">
              Enter only your name, even if your invitation is addressed to
              multiple people. You can RSVP for them in the next step!
            </HelperText>
            <Button
              onClick={() => findGuests(name)}
              disabled={loading}
              type="submit"
            >
              {loading ? <Spinner /> : 'Find my invitation'}
            </Button>
          </Card>
        </form>
      )}
      <GuestChooser guests={possibleGuests} />
    </section>
  )
}

function GuestChooser({ guests }: { guests: Guest[] }) {
  return (
    <pre>
      {guests.map((g) => (
        <div key={g.id}>{JSON.stringify(g, null, 2)}</div>
      ))}
    </pre>
  )
}

function InvitationCard({ guest }: { guest: Guest }) {
  return (
    <Card>
      <h3 className="text-xl font-semibold">{guest.name}</h3>
      <p>{guest.invitationId}</p>
    </Card>
  )
}
