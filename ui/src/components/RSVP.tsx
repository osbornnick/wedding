import { useState } from 'react'

export default function RSVP({
  initialSubmitted = false,
}: {
  initialSubmitted?: boolean
}) {
  const [submitted, setSubmitted] = useState(initialSubmitted)

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault()
    // TODO: wire up to API
    setSubmitted(true)
  }

  return (
    <section className="mx-auto max-w-2xl px-4 py-12">
      <h2 className="mb-4 text-center font-serif text-3xl font-semibold text-stone-800">
        RSVP
      </h2>
      <p className="mb-8 text-center text-stone-500">
        Please respond by <span className="font-medium">April 1, 2027</span>.
      </p>

      {submitted ? (
        <div className="rounded-2xl border border-green-200 bg-green-50 p-8 text-center">
          <p className="text-lg font-medium text-green-800">
            Thank you for your RSVP! We can't wait to celebrate with you.
          </p>
        </div>
      ) : (
        <form
          onSubmit={handleSubmit}
          className="space-y-5 rounded-2xl border border-stone-200 bg-stone-50 p-6 sm:p-8"
        >
          <div>
            <label
              htmlFor="name"
              className="mb-1 block text-sm font-medium text-stone-700"
            >
              Full Name
            </label>
            <input
              id="name"
              type="text"
              required
              placeholder="Jane Smith"
              className="w-full rounded-xl border border-stone-300 bg-white px-4 py-2.5 text-sm text-stone-800 placeholder-stone-400 focus:border-stone-500 focus:outline-none focus:ring-2 focus:ring-stone-200"
            />
          </div>

          <div>
            <label
              htmlFor="email"
              className="mb-1 block text-sm font-medium text-stone-700"
            >
              Email
            </label>
            <input
              id="email"
              type="email"
              required
              placeholder="jane@example.com"
              className="w-full rounded-xl border border-stone-300 bg-white px-4 py-2.5 text-sm text-stone-800 placeholder-stone-400 focus:border-stone-500 focus:outline-none focus:ring-2 focus:ring-stone-200"
            />
          </div>

          <div>
            <p className="mb-2 block text-sm font-medium text-stone-700">
              Will you attend?
            </p>
            <div className="flex gap-4">
              <label className="flex cursor-pointer items-center gap-2 text-sm text-stone-700">
                <input
                  type="radio"
                  name="attendance"
                  value="yes"
                  defaultChecked
                  className="accent-stone-800"
                />
                Joyfully accepts
              </label>
              <label className="flex cursor-pointer items-center gap-2 text-sm text-stone-700">
                <input
                  type="radio"
                  name="attendance"
                  value="no"
                  className="accent-stone-800"
                />
                Regretfully declines
              </label>
            </div>
          </div>

          <div>
            <label
              htmlFor="guests"
              className="mb-1 block text-sm font-medium text-stone-700"
            >
              Number of Guests (including yourself)
            </label>
            <select
              id="guests"
              className="w-full rounded-xl border border-stone-300 bg-white px-4 py-2.5 text-sm text-stone-800 focus:border-stone-500 focus:outline-none focus:ring-2 focus:ring-stone-200"
            >
              <option value="1">1</option>
              <option value="2">2</option>
              <option value="3">3</option>
              <option value="4">4</option>
            </select>
          </div>

          <div>
            <label
              htmlFor="dietary"
              className="mb-1 block text-sm font-medium text-stone-700"
            >
              Dietary Restrictions / Notes
            </label>
            <textarea
              id="dietary"
              rows={3}
              placeholder="Any dietary restrictions or special requests?"
              className="w-full rounded-xl border border-stone-300 bg-white px-4 py-2.5 text-sm text-stone-800 placeholder-stone-400 focus:border-stone-500 focus:outline-none focus:ring-2 focus:ring-stone-200"
            />
          </div>

          <button
            type="submit"
            className="w-full rounded-full bg-stone-800 px-5 py-3 text-sm font-semibold text-white transition hover:bg-stone-700 active:scale-[0.98]"
          >
            Send RSVP
          </button>
        </form>
      )}
    </section>
  )
}
