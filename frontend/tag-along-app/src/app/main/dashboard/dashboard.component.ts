import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AppService } from 'src/app/app.service';
import { Group } from 'src/app/models/group.model';
import Swal from 'sweetalert2'
import { finalize } from 'rxjs';
import { Event } from 'src/app/models/event.model';

@Component({
    selector: 'app-dashboard',
    templateUrl: './dashboard.component.html',
})
export class DashboardComponent implements OnInit {

    isGroupLoading = false;
    groups: Group[] = [];

    isEventLoading = false;
    events: Event[] = [];

    constructor(
        private readonly http: HttpClient,
        public readonly appService: AppService
    ) { }

    ngOnInit(): void {
        this.isGroupLoading = true;
        this.http.get<Group[]>('/groups/user/' + this.appService.loggedInUser?.id)
            .pipe(finalize(() => this.isGroupLoading = false))
            .subscribe({
                next: data => this.groups = data,
                error: err => Swal.fire({
                    icon: 'error',
                    title: 'Internal Server Error',
                    text: (err && err.error) ? err.error.error : 'Something went wrong fetching group details',
                })
            });

        this.isEventLoading = true;
        this.http.get<Event[]>('/events/user/' + this.appService.loggedInUser?.id)
            .pipe(finalize(() => this.isEventLoading = false))
            .subscribe({
                next: data => this.events = data,
                error: err => Swal.fire({
                    icon: 'error',
                    title: 'Internal Server Error',
                    text: (err && err.error) ? err.error.error : 'Something went wrong fetching group details',
                })
            });
    }


    isEventAccepted(event: Event): boolean {
        return event.acceptedUsers.includes(this.appService.loggedInUser!.id)
    }

}
