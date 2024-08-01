import { Component, inject } from "@angular/core";
import { AsyncPipe, CommonModule } from "@angular/common";
import { SubscribersService } from "../../services";
import { Observable, tap } from "rxjs";
import { EmptySubscriber, Subscriber } from "../../models";

@Component({
  selector: "app-subscriber-container",
  standalone: true,
  imports: [CommonModule, AsyncPipe],
  templateUrl: "./subscriber-container.component.html",
  styleUrl: "./subscriber-container.component.css",
})
export class SubscriberContainerComponent {
  private subscriberService = inject(SubscribersService);
  winner: Subscriber = EmptySubscriber

  subscribers$: Observable<Subscriber[]> =
    this.subscriberService.getAllSubscribers();

  getWinner() {
    this.winner = this.subscriberService.getWinner();
  }
}
