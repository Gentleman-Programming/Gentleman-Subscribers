import { Component } from "@angular/core";
import { RouterModule } from "@angular/router";
import { NxWelcomeComponent } from "./nx-welcome.component";
import { SubscriberContainerComponent } from "../components";

@Component({
  standalone: true,
  imports: [NxWelcomeComponent, RouterModule, SubscriberContainerComponent],
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrl: "./app.component.scss",
})
export class AppComponent {
  title = "gentleman-subscribers";
}
