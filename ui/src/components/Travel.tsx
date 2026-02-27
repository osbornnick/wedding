export default function Travel() {
  return (
    <section className="mx-auto max-w-2xl px-4 py-12">
      <h2 className="mb-8 text-center font-serif text-3xl font-semibold text-stone-800">
        Travel &amp; Accommodations
      </h2>

      <div className="space-y-8">
        <div className="rounded-2xl border border-stone-200 bg-stone-50 p-6">
          <h3 className="mb-2 text-lg font-semibold text-stone-800">Venue</h3>
          <p className="text-stone-600">
            <span className="font-medium">TBD Venue Name</span>
            <br />
            123 Celebration Lane
            <br />
            City, State 00000
          </p>
        </div>

        <div className="rounded-2xl border border-stone-200 bg-stone-50 p-6">
          <h3 className="mb-2 text-lg font-semibold text-stone-800">
            Hotel Block
          </h3>
          <p className="text-stone-600 mb-3">
            We have reserved a room block at the hotel below. Use the code{' '}
            <span className="font-semibold text-stone-800">NICKYSARAH</span>{' '}
            when booking.
          </p>
          <p className="text-stone-600">
            <span className="font-medium">TBD Hotel Name</span>
            <br />
            456 Stay Avenue
            <br />
            City, State 00000
            <br />
            (555) 000-0000
          </p>
          <a
            href="#"
            className="mt-4 inline-block rounded-full bg-stone-800 px-5 py-2.5 text-sm font-medium text-white no-underline transition hover:bg-stone-700"
          >
            Book Your Room
          </a>
        </div>

        <div className="rounded-2xl border border-stone-200 bg-stone-50 p-6">
          <h3 className="mb-2 text-lg font-semibold text-stone-800">
            Getting Here
          </h3>
          <p className="text-stone-600">
            The venue is approximately 30 minutes from the nearest airport. We
            recommend renting a car or using a rideshare service. Parking is
            available on-site at no charge.
          </p>
        </div>
      </div>
    </section>
  )
}
