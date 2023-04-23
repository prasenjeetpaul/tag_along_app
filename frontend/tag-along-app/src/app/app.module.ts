import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FooterComponent } from './footer/footer.component';
import { HeaderComponent } from './header/header.component';
import { MainComponent } from './main/main.component';
import { AppService } from './app.service';
import { LoginComponent } from './main/login/login.component';
import { RegisterComponent } from './main/register/register.component';
import { DashboardComponent } from './main/dashboard/dashboard.component';
import { GroupDetailComponent } from './main/group/group-details/group-detail.component';
import { NewGroupComponent } from './main/group/new-group/new-group.component';
import { NewEventComponent } from './main/event/new-event/new-event.component';
import { AuthGuardService } from './auth.guard';
import { ReactiveFormsModule } from '@angular/forms';
import { UserPillComponent } from './util/user-pill/user-pill.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    MainComponent,
    FooterComponent,
    LoginComponent,
    RegisterComponent,
    DashboardComponent,
    GroupDetailComponent,
    NewGroupComponent,
    NewEventComponent,
    UserPillComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [
    AppService,
    AuthGuardService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
