package main

import (
	"time"
)

type Message struct {
	ID        uint      `json:"id"`
	Subject   string    `json:"subject"`
	Sender    string    `json:"sender"`
	Recipient string    `json:"recipient"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

var messages = []Message{
	{
		ID:        1,
		Subject:   "Welcome to Nova Laboratories Internal Messaging System!",
		Sender:    "a.johnson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Dear Team,

I am excited to announce the launch of our new internal messaging system at Nova Laboratories! This platform is designed to enhance our communication and collaboration.

You can now send messages directly to individuals or to distribution lists. For example, you can use "all@" to send a message to everyone in the company.

Please feel free to start conversations, share updates, and participate in discussions. Your feedback is valuable as we continue to improve this system.

Best regards,
Alex Johnson
Head of IT`,
		Timestamp: time.Date(1985, time.March, 5, 9, 0, 0, 0, time.UTC),
	},
	{
		ID:        2,
		Subject:   "Hello Everyone!",
		Sender:    "e.carter@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

I hope you're all doing well. I'm excited to see how this new messaging system can improve our communication.

Best,
Dr. Evelyn Carter
Head of Research`,
		Timestamp: time.Date(1985, time.March, 5, 10, 0, 0, 0, time.UTC),
	},
	{
		ID:        3,
		Subject:   "Greetings!",
		Sender:    "c.thompson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello team,

Looking forward to more streamlined communication with this new system. Let's make the most of it!

Best,
Clara Thompson
Finance Manager`,
		Timestamp: time.Date(1985, time.March, 5, 11, 0, 0, 0, time.UTC),
	},
	{
		ID:        4,
		Subject:   "Hi Team!",
		Sender:    "d.lee@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

Just wanted to drop a quick message to say hello and that I'm here to help with any HR-related questions.

Cheers,
David Lee
HR Director`,
		Timestamp: time.Date(1985, time.March, 5, 12, 0, 0, 0, time.UTC),
	},
	{
		ID:        5,
		Subject:   "Excited to Connect!",
		Sender:    "e.davis@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello all,

I'm really excited about this new system. It's going to be great for our projects.

Best regards,
Emily Davis
Lead Developer`,
		Timestamp: time.Date(1985, time.March, 5, 13, 0, 0, 0, time.UTC),
	},
	{
		ID:        6,
		Subject:   "password123",
		Sender:    "g.anderson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message:   `password123`,
		Timestamp: time.Date(1985, time.March, 5, 14, 0, 0, 0, time.UTC),
	},
	{
		ID:        7,
		Subject:   "Hello Everyone!",
		Sender:    "f.wilson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

Looking forward to collaborating more effectively with all of you through this new platform.

Best,
Frank Wilson
Marketing Manager`,
		Timestamp: time.Date(1985, time.March, 5, 15, 0, 0, 0, time.UTC),
	},
	{
		ID:        8,
		Subject:   "Hi Everyone!",
		Sender:    "h.brown@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello everyone,

Excited to see how this new system will improve our operations.

Best,
Henry Brown
Operations Manager`,
		Timestamp: time.Date(1985, time.March, 5, 16, 0, 0, 0, time.UTC),
	},
	{
		ID:        9,
		Subject:   "Welcome!",
		Sender:    "g.martinez@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

Welcome to the new messaging system! If you have any questions or need assistance, feel free to reach out.

Best regards,
Grace Martinez
Administrative Assistant`,
		Timestamp: time.Date(1985, time.March, 5, 17, 0, 0, 0, time.UTC),
	},
	{
		ID:        10,
		Subject:   "Hello All!",
		Sender:    "i.white@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

Looking forward to ensuring the quality of our products with this new system in place.

Best,
Irene White
Quality Assurance Lead`,
		Timestamp: time.Date(1985, time.March, 5, 18, 0, 0, 0, time.UTC),
	},
	{
		ID:        11,
		Subject:   "Greetings!",
		Sender:    "j.green@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello team,

Excited to see how we can leverage this new system for our robotics projects.

Best,
Jack Green
Robotics Engineer`,
		Timestamp: time.Date(1985, time.March, 5, 19, 0, 0, 0, time.UTC),
	},
	{
		ID:        12,
		Subject:   "Hi Everyone!",
		Sender:    "k.roberts@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi all,

Can't wait to analyze the data we'll gather through this new messaging platform.

Best,
Kelly Roberts
Data Analyst`,
		Timestamp: time.Date(1985, time.March, 5, 20, 0, 0, 0, time.UTC),
	},
	{
		ID:        13,
		Subject:   "Password Reset Notice",
		Sender:    "a.johnson@nova-messages.mentats.org",
		Recipient: "g.anderson@nova-messages.mentats.org",
		Message: `Hi Greg,

Your password has been reset due to a security policy violation. Please check your email for the new temporary password and make sure to change it immediately. Don't forget to read our internal policies on password usage.

Best regards,
Alex Johnson
Head of IT`,
		Timestamp: time.Date(1985, time.March, 6, 9, 0, 0, 0, time.UTC),
	},
	{
		ID:        14,
		Subject:   "Hello Everyone!",
		Sender:    "l.king@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

Looking forward to improving our network infrastructure with this new messaging system.

Best,
Liam King
Network Administrator`,
		Timestamp: time.Date(1985, time.March, 6, 10, 0, 0, 0, time.UTC),
	},
	{
		ID:        15,
		Subject:   "Welcome!",
		Sender:    "m.taylor@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello everyone,

Excited to manage our product development cycles more effectively with this new tool.

Best,
Monica Taylor
Product Manager`,
		Timestamp: time.Date(1985, time.March, 6, 11, 0, 0, 0, time.UTC),
	},
	{
		ID:        16,
		Subject:   "Hi All!",
		Sender:    "n.scott@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

Let's work together to ensure our security measures are up to date with this new system.

Best,
Nathan Scott
Security Specialist`,
		Timestamp: time.Date(1985, time.March, 6, 12, 0, 0, 0, time.UTC),
	},
	{
		ID:        17,
		Subject:   "Hi Everyone!",
		Sender:    "f.wilson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello team,

Excited to collaborate on new marketing strategies with this new platform.

Best,
Frank Wilson
Marketing Manager`,
		Timestamp: time.Date(1985, time.March, 6, 13, 0, 0, 0, time.UTC),
	},
	{
		ID:        18,
		Subject:   "Hello!",
		Sender:    "h.brown@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

Looking forward to optimizing our operations with the help of this new system.

Best,
Henry Brown
Operations Manager`,
		Timestamp: time.Date(1985, time.March, 6, 14, 0, 0, 0, time.UTC),
	},
	{
		ID:        19,
		Subject:   "Welcome!",
		Sender:    "g.martinez@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello team,

Remember, if you have any questions or need help, I'm here for you.

Best regards,
Grace Martinez
Administrative Assistant`,
		Timestamp: time.Date(1985, time.March, 6, 15, 0, 0, 0, time.UTC),
	},
	{
		ID:        20,
		Subject:   "Team Outing",
		Sender:    "m.taylor@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

We are planning a team outing next month. Please let me know your availability and any suggestions for activities.

Best,
Monica Taylor
Product Manager`,
		Timestamp: time.Date(1985, time.March, 7, 9, 0, 0, 0, time.UTC),
	},
	{
		ID:        21,
		Subject:   "Found Floppy Disk",
		Sender:    "g.anderson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi all,

I found a floppy disk in the parking lot, but it didn't have anything interesting on it. Does anyone know whose it is?

Best,
Greg Anderson
Senior Engineer`,
		Timestamp: time.Date(1985, time.March, 7, 10, 0, 0, 0, time.UTC),
	},
	{
		ID:        22,
		Subject:   "New Research Initiative",
		Sender:    "e.carter@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello team,

We are launching a new research initiative focused on advanced AI applications. If you are interested in joining, please contact me.

Best,
Dr. Evelyn Carter
Head of Research`,
		Timestamp: time.Date(1985, time.March, 7, 11, 0, 0, 0, time.UTC),
	},
	{
		ID:        23,
		Subject:   "Finance Department Meeting",
		Sender:    "c.thompson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

The finance department will have a meeting on Friday at 2 PM to discuss the quarterly budget. Please be on time.

Best,
Clara Thompson
Finance Manager`,
		Timestamp: time.Date(1985, time.March, 7, 12, 0, 0, 0, time.UTC),
	},
	{
		ID:        24,
		Subject:   "HR Policy Update",
		Sender:    "d.lee@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Dear all,

Please review the updated HR policies that have been posted on the intranet. If you have any questions, feel free to contact me.

Best,
David Lee
HR Director`,
		Timestamp: time.Date(1985, time.March, 7, 13, 0, 0, 0, time.UTC),
	},
	{
		ID:        25,
		Subject:   "Client Visit",
		Sender:    "f.wilson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

We will have clients visiting our office next week. Please ensure your workspaces are tidy and be prepared to showcase your projects.

Best,
Frank Wilson
Marketing Manager`,
		Timestamp: time.Date(1985, time.March, 7, 14, 0, 0, 0, time.UTC),
	},
	{
		ID:        26,
		Subject:   "Networking Event",
		Sender:    "h.brown@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello everyone,

There is a networking event scheduled for next Thursday evening. Please RSVP if you plan to attend.

Best,
Henry Brown
Operations Manager`,
		Timestamp: time.Date(1985, time.March, 7, 15, 0, 0, 0, time.UTC),
	},
	{
		ID:        27,
		Subject:   "Help Needed with Desktop",
		Sender:    "g.anderson@nova-messages.mentats.org",
		Recipient: "it@nova-messages.mentats.org",
		Message: `Hi IT,

My desktop computer has been acting strangely lately. Can someone please come and take a look at it?

Thanks,
Greg Anderson
Senior Engineer`,
		Timestamp: time.Date(1985, time.March, 7, 16, 0, 0, 0, time.UTC),
	},
	{
		ID:        28,
		Subject:   "Project Deadline",
		Sender:    "e.davis@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

Just a reminder that the deadline for the current project is next Monday. Please ensure all tasks are completed by then.

Best regards,
Emily Davis
Lead Developer`,
		Timestamp: time.Date(1985, time.March, 7, 17, 0, 0, 0, time.UTC),
	},
	{
		ID:        29,
		Subject:   "Security Reminder",
		Sender:    "n.scott@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Dear colleagues,

This is a reminder about our basic security measures. Please ensure you lock your workstations when unattended, use strong passwords, and follow the procedure for requisitioning new floppy disks. If you have any questions, please refer to the security policy document or contact the IT department.

Best regards,
Nathan Scott
Security Specialist`,
		Timestamp: time.Date(1985, time.March, 7, 18, 0, 0, 0, time.UTC),
	},
	{
		ID:        30,
		Subject:   "St. Patrick's Day Party Planning",
		Sender:    "g.martinez@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

It's time to start planning our annual St. Patrick's Day party! If you have any ideas for activities, decorations, or food, please let me know.

Best regards,
Grace Martinez
Administrative Assistant`,
		Timestamp: time.Date(1985, time.March, 8, 9, 0, 0, 0, time.UTC),
	},
	{
		ID:        31,
		Subject:   "Weekly Status Report",
		Sender:    "e.carter@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello team,

Please submit your weekly status reports by end of day today.

Best,
Dr. Evelyn Carter
Head of Research`,
		Timestamp: time.Date(1985, time.March, 8, 10, 0, 0, 0, time.UTC),
	},
	{
		ID:        32,
		Subject:   "St. Patrick's Day Costume Contest",
		Sender:    "f.wilson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi all,

As part of our St. Patrick's Day celebration, we will be having a costume contest. Start planning your outfits, and may the best costume win!

Best,
Frank Wilson
Marketing Manager`,
		Timestamp: time.Date(1985, time.March, 8, 11, 0, 0, 0, time.UTC),
	},
	{
		ID:        33,
		Subject:   "Project Update Meeting",
		Sender:    "e.davis@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

There will be a project update meeting tomorrow at 2 PM. Please come prepared to discuss your progress.

Best regards,
Emily Davis
Lead Developer`,
		Timestamp: time.Date(1985, time.March, 8, 12, 0, 0, 0, time.UTC),
	},
	{
		ID:        34,
		Subject:   "St. Patrick's Day Party Menu",
		Sender:    "h.brown@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello everyone,

We are working on the menu for the St. Patrick's Day party. If you have any food or drink requests, please let us know by the end of the week.

Best,
Henry Brown
Operations Manager`,
		Timestamp: time.Date(1985, time.March, 8, 13, 0, 0, 0, time.UTC),
	},
	{
		ID:        35,
		Subject:   "Maintenance Request",
		Sender:    "j.green@nova-messages.mentats.org",
		Recipient: "it@nova-messages.mentats.org",
		Message: `Hi IT,

I need some assistance with the robotics lab equipment. Could someone please check the systems?

Thanks,
Jack Green
Robotics Engineer`,
		Timestamp: time.Date(1985, time.March, 8, 14, 0, 0, 0, time.UTC),
	},
	{
		ID:        36,
		Subject:   "Client Presentation",
		Sender:    "m.taylor@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

We have a client presentation scheduled for next Monday. Please ensure all materials are ready and submitted for review by Friday.

Best,
Monica Taylor
Product Manager`,
		Timestamp: time.Date(1985, time.March, 8, 15, 0, 0, 0, time.UTC),
	},
	{
		ID:        37,
		Subject:   "Lunch and Learn Session",
		Sender:    "n.scott@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Dear colleagues,

We will be having a lunch and learn session on cybersecurity next Wednesday. Please join us in the conference room at noon.

Best regards,
Nathan Scott
Security Specialist`,
		Timestamp: time.Date(1985, time.March, 8, 16, 0, 0, 0, time.UTC),
	},
	{
		ID:        38,
		Subject:   "Office Supplies",
		Sender:    "g.martinez@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

We will be placing an order for office supplies this Friday. If you need anything specific, please let me know by Thursday.

Best regards,
Grace Martinez
Administrative Assistant`,
		Timestamp: time.Date(1985, time.March, 8, 17, 0, 0, 0, time.UTC),
	},
	{
		ID:        39,
		Subject:   "St. Patrick's Day Decoration Volunteers",
		Sender:    "d.lee@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

We need volunteers to help with decorating the office for the St. Patrick's Day party. If you are available, please reply to this message.

Cheers,
David Lee
HR Director`,
		Timestamp: time.Date(1985, time.March, 8, 18, 0, 0, 0, time.UTC),
	},
	{
		ID:        40,
		Subject:   "Monthly Performance Review",
		Sender:    "c.thompson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello team,

This is a reminder that monthly performance reviews are due next Friday. Please make sure to complete your self-evaluations and submit them to your managers.

Best,
Clara Thompson
Finance Manager`,
		Timestamp: time.Date(1985, time.March, 9, 9, 0, 0, 0, time.UTC),
	},
	{
		ID:        41,
		Subject:   "New Onboarding System Announcement",
		Sender:    "d.lee@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Dear all,

We are excited to announce the launch of our new onboarding system for new employees. This system is designed to streamline the onboarding process and ensure new hires have all the resources they need to get started quickly.

Please review the onboarding documentation available on the intranet and provide any feedback you may have.

Best,
David Lee
HR Director`,
		Timestamp: time.Date(1985, time.March, 9, 10, 0, 0, 0, time.UTC),
	},
	{
		ID:        42,
		Subject:   "St. Patrick's Day Party Reminder",
		Sender:    "g.martinez@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

Just a reminder that the St. Patrick's Day party is next week. We are still looking for volunteers to help with the decorations. Let me know if you can assist.

Best regards,
Grace Martinez
Administrative Assistant`,
		Timestamp: time.Date(1985, time.March, 9, 11, 0, 0, 0, time.UTC),
	},
	{
		ID:        43,
		Subject:   "System Maintenance Notification",
		Sender:    "a.johnson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Dear colleagues,

Please be advised that there will be a system maintenance window this Saturday from 10 PM to 4 AM. During this time, the internal messaging system may be unavailable.

Thank you for your understanding.

Best regards,
Alex Johnson
Head of IT`,
		Timestamp: time.Date(1985, time.March, 9, 12, 0, 0, 0, time.UTC),
	},
	{
		ID:        44,
		Subject:   "Research Grant Application",
		Sender:    "e.carter@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

We have an opportunity to apply for a new research grant. If you have any project proposals, please submit them to me by the end of the week.

Best,
Dr. Evelyn Carter
Head of Research`,
		Timestamp: time.Date(1985, time.March, 9, 13, 0, 0, 0, time.UTC),
	},
	{
		ID:        45,
		Subject:   "Office Clean-Up Day",
		Sender:    "h.brown@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello everyone,

We will have an office clean-up day this Friday. Please take some time to tidy up your workspaces and common areas.

Best,
Henry Brown
Operations Manager`,
		Timestamp: time.Date(1985, time.March, 9, 14, 0, 0, 0, time.UTC),
	},
	{
		ID:        46,
		Subject:   "Marketing Campaign Launch",
		Sender:    "f.wilson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

Our new marketing campaign will launch next Monday. Please review the campaign materials and be prepared to support any customer inquiries.

Best,
Frank Wilson
Marketing Manager`,
		Timestamp: time.Date(1985, time.March, 9, 15, 0, 0, 0, time.UTC),
	},
	{
		ID:        47,
		Subject:   "IT Equipment Inventory",
		Sender:    "l.king@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi all,

We will be conducting an inventory of all IT equipment next week. Please make sure your devices are labeled and accounted for.

Best,
Liam King
Network Administrator`,
		Timestamp: time.Date(1985, time.March, 9, 16, 0, 0, 0, time.UTC),
	},
	{
		ID:        48,
		Subject:   "Software Development Workshop",
		Sender:    "e.davis@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

We are organizing a software development workshop next Thursday. If you are interested in attending, please sign up by Tuesday.

Best regards,
Emily Davis
Lead Developer`,
		Timestamp: time.Date(1985, time.March, 9, 17, 0, 0, 0, time.UTC),
	},
	{
		ID:        49,
		Subject:   "Company Picnic Planning",
		Sender:    "m.taylor@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

We are starting to plan this year's company picnic. If you have any suggestions for locations or activities, please let me know.

Best,
Monica Taylor
Product Manager`,
		Timestamp: time.Date(1985, time.March, 9, 18, 0, 0, 0, time.UTC),
	},
	{
		ID:        50,
		Subject:   "Team Lunch",
		Sender:    "g.martinez@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

We are planning a team lunch next Wednesday at 1 PM. Please let me know if you have any dietary restrictions.

Best regards,
Grace Martinez
Administrative Assistant`,
		Timestamp: time.Date(1985, time.March, 10, 9, 0, 0, 0, time.UTC),
	},
	{
		ID:        51,
		Subject:   "Welcome Aboard!",
		Sender:    "d.lee@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

Please join me in welcoming our new colleague, Sarah Johnson, to the team! Sarah will be joining the Research Department as a Junior Researcher. We are excited to have her on board.

Best,
David Lee
HR Director`,
		Timestamp: time.Date(1985, time.March, 10, 10, 0, 0, 0, time.UTC),
	},
	{
		ID:        52,
		Subject:   "New Research Equipment",
		Sender:    "e.carter@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello team,

We have received new equipment for the research lab. Please schedule time with the lab manager if you need to use it.

Best,
Dr. Evelyn Carter
Head of Research`,
		Timestamp: time.Date(1985, time.March, 10, 11, 0, 0, 0, time.UTC),
	},
	{
		ID:        53,
		Subject:   "Quarterly Budget Review",
		Sender:    "c.thompson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

This is a reminder that the quarterly budget review meeting is scheduled for next Monday. Please prepare your reports and submit them by Friday.

Best,
Clara Thompson
Finance Manager`,
		Timestamp: time.Date(1985, time.March, 10, 12, 0, 0, 0, time.UTC),
	},
	{
		ID:        54,
		Subject:   "Bug Report: Onboarding System Issue",
		Sender:    "a.johnson@nova-messages.mentats.org",
		Recipient: "it@nova-messages.mentats.org",
		Message: `Hi IT team,

We've identified a bug in the new onboarding system where it incorrectly created a new employee's account with superadmin privileges. Please investigate and resolve this issue as soon as possible.

Best regards,
Alex Johnson
Head of IT`,
		Timestamp: time.Date(1985, time.March, 10, 13, 0, 0, 0, time.UTC),
	},
	{
		ID:        55,
		Subject:   "Marketing Strategy Meeting",
		Sender:    "f.wilson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

There will be a marketing strategy meeting this Friday at 2 PM. Please come prepared with your ideas and suggestions.

Best,
Frank Wilson
Marketing Manager`,
		Timestamp: time.Date(1985, time.March, 10, 14, 0, 0, 0, time.UTC),
	},
	{
		ID:        56,
		Subject:   "Robotics Workshop",
		Sender:    "j.green@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

We are organizing a robotics workshop next Tuesday at 10 AM. If you are interested in attending, please sign up by Monday.

Best,
Jack Green
Robotics Engineer`,
		Timestamp: time.Date(1985, time.March, 10, 15, 0, 0, 0, time.UTC),
	},
	{
		ID:        57,
		Subject:   "Account Permissions Update",
		Sender:    "a.johnson@nova-messages.mentats.org",
		Recipient: "s.johnson@nova-messages.mentats.org",
		Message: `Hi Sarah,

Your account permissions have been updated to the correct level, and the superadmin access has been removed. A bugfix for the onboarding system has been added to the backlog.

Best regards,
Alex Johnson
Head of IT`,
		Timestamp: time.Date(1985, time.March, 10, 16, 0, 0, 0, time.UTC),
	},
	{
		ID:        58,
		Subject:   "Team Building Activity",
		Sender:    "d.lee@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

We are planning a team-building activity next Thursday afternoon. Please RSVP if you plan to attend.

Best,
David Lee
HR Director`,
		Timestamp: time.Date(1985, time.March, 10, 17, 0, 0, 0, time.UTC),
	},
	{
		ID:        59,
		Subject:   "Office Renovation Update",
		Sender:    "h.brown@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello everyone,

The office renovation is progressing well. Please be aware that the west wing will be closed for the next two weeks for painting and flooring work.

Best,
Henry Brown
Operations Manager`,
		Timestamp: time.Date(1985, time.March, 10, 18, 0, 0, 0, time.UTC),
	},
	{
		ID:        60,
		Subject:   "Project Deadline Reminder",
		Sender:    "e.davis@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

Just a reminder that the project deadline is this Friday. Please ensure all deliverables are completed and submitted.

Best regards,
Emily Davis
Lead Developer`,
		Timestamp: time.Date(1985, time.March, 15, 9, 0, 0, 0, time.UTC),
	},
	{
		ID:        61,
		Subject:   "Office Supplies Order",
		Sender:    "g.martinez@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

We will be placing an order for office supplies tomorrow. If you need anything specific, please let me know by the end of the day.

Best regards,
Grace Martinez
Administrative Assistant`,
		Timestamp: time.Date(1985, time.March, 15, 10, 0, 0, 0, time.UTC),
	},
	{
		ID:        62,
		Subject:   "Client Feedback Session",
		Sender:    "f.wilson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

There will be a client feedback session this Thursday at 3 PM. Please attend if you worked on the recent projects.

Best,
Frank Wilson
Marketing Manager`,
		Timestamp: time.Date(1985, time.March, 15, 11, 0, 0, 0, time.UTC),
	},
	{
		ID:        63,
		Subject:   "Network Maintenance",
		Sender:    "l.king@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi all,

Please be informed that there will be network maintenance this Saturday from 10 PM to 2 AM. The network may be unavailable during this time.

Best,
Liam King
Network Administrator`,
		Timestamp: time.Date(1985, time.March, 15, 12, 0, 0, 0, time.UTC),
	},
	{
		ID:        64,
		Subject:   "Team Outing",
		Sender:    "h.brown@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hello everyone,

We are planning a team outing for the last Friday of the month. Please let me know if you have any suggestions for activities.

Best,
Henry Brown
Operations Manager`,
		Timestamp: time.Date(1985, time.March, 15, 13, 0, 0, 0, time.UTC),
	},
	{
		ID:        65,
		Subject:   "St. Patrick's Day Party Update",
		Sender:    "g.martinez@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

The St. Patrick's Day party is coming up this weekend! We have planned some fun activities and hope to see you all there.

Best regards,
Grace Martinez
Administrative Assistant`,
		Timestamp: time.Date(1985, time.March, 15, 14, 0, 0, 0, time.UTC),
	},
	{
		ID:        66,
		Subject:   "Research Collaboration Opportunity",
		Sender:    "e.carter@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi team,

We have an exciting opportunity to collaborate on a research project with another institution. If you are interested, please reach out to me for more details.

Best,
Dr. Evelyn Carter
Head of Research`,
		Timestamp: time.Date(1985, time.March, 15, 15, 0, 0, 0, time.UTC),
	},
	{
		ID:        67,
		Subject:   "Bug Fix Release",
		Sender:    "a.johnson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Dear colleagues,

We will be releasing a bug fix for the internal messaging system tonight at 8 PM. Please save your work and log out before then.

Best regards,
Alex Johnson
Head of IT`,
		Timestamp: time.Date(1985, time.March, 15, 16, 0, 0, 0, time.UTC),
	},
	{
		ID:        68,
		Subject:   "New Office Policies",
		Sender:    "d.lee@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Dear all,

Please review the updated office policies that have been posted on the intranet. If you have any questions, feel free to contact me.

Best,
David Lee
HR Director`,
		Timestamp: time.Date(1985, time.March, 15, 17, 0, 0, 0, time.UTC),
	},
	{
		ID:        69,
		Subject:   "Monthly Financial Report",
		Sender:    "c.thompson@nova-messages.mentats.org",
		Recipient: "all@nova-messages.mentats.org",
		Message: `Hi everyone,

The monthly financial report is now available on the shared drive. Please review it and let me know if you have any questions.

Best,
Clara Thompson
Finance Manager`,
		Timestamp: time.Date(1985, time.March, 15, 18, 0, 0, 0, time.UTC),
	},
}
