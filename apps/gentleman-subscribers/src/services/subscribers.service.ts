import { HttpClient } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { Subscriber } from "../models";
import { Observable, tap } from "rxjs";

@Injectable({
  providedIn: "root",
})
export class SubscribersService {
  private mailUrl = "http://0.0.0.0:3000/apis/v1/getAllSubscribers";
  private http = inject(HttpClient);
  private subscribers: Subscriber[] = [];

  getAllSubscribers(): Observable<Subscriber[]> {
    return this.http.get<Subscriber[]>(this.mailUrl).pipe(tap(subs => this.subscribers = subs));
  }

  getWinner(): Subscriber {
    if (this.subscribers.length === 0) {
      throw new Error("el arreglo está vacio");
    }

    const randomIndex = Math.floor(Math.random() * this.subscribers.length)

    return this.subscribers[randomIndex];
  }
}
