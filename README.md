# LIFT Corporate Web Functional Specification

## Overview

This project is a web portal designed for corporate use, which will be white-labeled with the company's branding, color theme, and logo. The portal is built to be responsive, ensuring compatibility with various devices, including mobile browsers. User authentication is secured through OTP verification during the login process.

## Features

### 1. Main Menu

#### 1.1 Profile
- Display corporate profile information (name, ID, etc.).
- Admin user list includes First name, Last name, email id, accounts email.
- Display credit limit and payment due date, both weekly and monthly.
- Show discount percentages and caps.

#### 1.2 Notifications
- Display notifications related to promotions and onboarding requests.

#### 1.3 Users
- List of users with assigned groups.
- Ability to filter users and change their groups.
- Admin has the option to place special bookings on behalf of users, overriding group limitations.

#### 1.4 Groups
- List of groups with the ability for admin users to edit group details.

#### 1.5 Reports
- Display various reports (details specified in section 3).

#### 1.6 Audit Logs
- Showcase audit records (details specified in section 4).

#### 1.7 Contact Us
- A page with various options to contact the Lift Team.
- Includes a contact form to submit custom messages to the Lift Team.
- An email will be triggered to the support email for the Lift support email, including a unique ID.

## Additional Notes

- This project uses Go as the primary programming language.
- Ensure proper configuration, including database settings, email configurations, etc.
- For more details on specific sections, please refer to the corresponding sections below.

## Section Details

### 2. Admin User List (Profile)

...

### 3. Reports

...

### 4. Audit Logs

...

## License

This project is licensed under the [License Name] - see the [LICENSE.md](LICENSE.md) file for details.
