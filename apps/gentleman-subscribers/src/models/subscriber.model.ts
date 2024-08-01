export interface Subscriber {
  name: string;
  email: string;
  id: string;
  country: string;
}

export const EmptySubscriber: Subscriber = {
  name: "",
  email: "",
  id: "",
  country: ""
}
