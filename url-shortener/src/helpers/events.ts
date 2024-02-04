import { EventName, storeEvent } from "../services/events";


class EventsSDK {
  private static instance: EventsSDK;
  private events: string[];

  private constructor() {
    this.events = [];
  }

  public static getInstance(): EventsSDK {
    if (!EventsSDK.instance) {
      EventsSDK.instance = new EventsSDK();
    }
    return EventsSDK.instance;
  }

  public handleEvent(event: EventName): void {
    storeEvent(event)
  }
}

const eventsSDK = EventsSDK.getInstance();

document.addEventListener("click", () => {
  eventsSDK.handleEvent(EventName.Click);
});
