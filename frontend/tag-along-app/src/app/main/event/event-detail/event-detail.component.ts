import { Component, Input, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AppService } from 'src/app/app.service';
import Swal from 'sweetalert2'
import { Event } from 'src/app/models/event.model';

@Component({
    selector: 'app-event-detail',
    templateUrl: './event-detail.component.html',
})
export class EventDetailComponent implements OnInit {

    @Input() event!: Event;
    isEventAccepted: boolean = false;

    constructor(
        private readonly http: HttpClient,
        public readonly appService: AppService
    ) { }

    ngOnInit(): void {
        this.isEventAccepted = this.event.acceptedUsers.includes(this.appService.loggedInUser!.id)
    }

}
