import { Carousel } from 'flowbite-react'

const carouselTheme = {
  control: {
    base: 'inline-flex h-10 w-10 items-center justify-center rounded-full bg-stone-800/60 group-hover:bg-stone-800/80 dark:bg-white/20 dark:group-hover:bg-white/40 sm:h-12 sm:w-12',
    icon: 'h-5 w-5 text-white dark:text-white sm:h-6 sm:w-6',
  },
}

export default function PhotoCarousel() {
  return (
    <div className="m-5 h-[75vh] max-h-[480px]">
      <Carousel className="w-full" indicators={false} theme={carouselTheme}>
        <CarouselItem
          src="/images/engagement.jpg"
          alt="Nicky and Sarah"
          description="Dec 20 2025, Sarah and Nicky returned to Longfellow park, the site of their first date 7 years prior, and got engaged!"
        />
        <CarouselItem
          src="/images/davincis.jpg"
          alt="Da Vinci's"
          description="Sep 30, 2018, a thank you selfie for our friend whose car we borrowed for a dinner date at DaVinci's, Lewiston's premier spot for garlic knots."
        />
        <CarouselItem
          src="/images/babies.jpg"
          alt="young Nicky and Sarah"
          description="Mar 23, 2019, young Nicky and Sarah at the 'Gala' dance at Bates. Nicky was a senior soon to graduate, and Sarah a sophomore. Truly so baby."
        />
        <CarouselItem
          src="/images/copenhagen.jpg"
          alt="Copenhagen"
          description="Oct 28, 2019, Nicky visited Sarah in Copenhagen, where she was studying abroad. He got sick and they learned they can do hard things."
        />
        <CarouselItem
          src="/images/gw.jpg"
          alt="accepted to GW!"
          description="Mar 21, 2024, Sarah accepted her offer to George Washington University while we were at Logan airport! This solidified our next chapter together in DC."
        />
        <CarouselItem
          src="/images/zion.jpg"
          alt="Zion National Park"
          description="May 30, 2024, hiking The Narrows at Zion National Park with Sarah's family. The matching jersey's are from Cold Front, Sarah's frisbee team at Bates. Nicky is a big fan."
        />
      </Carousel>
    </div>
  )
}

function CarouselItem({
  src,
  alt,
  description,
}: {
  src: string
  alt: string
  description: string
}) {
  return (
    <div className="h-full flex flex-col items-center justify-center overflow-hidden">
      <img src={src} alt={alt} className="min-h-0 object-contain rounded-xl" />
      <p className="mt-2 text-center text-sm text-stone-500 dark:text-stone-300">
        {description}
      </p>
    </div>
  )
}
