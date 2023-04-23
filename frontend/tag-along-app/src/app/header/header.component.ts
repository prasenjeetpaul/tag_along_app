import { Component } from '@angular/core';
import { AppService } from '../app.service';
import { Router } from '@angular/router';

@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
})
export class HeaderComponent {
    constructor(
        private readonly router: Router,
        public readonly appService: AppService,
    ) { }
}

