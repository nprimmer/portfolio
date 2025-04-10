# Crafting The Challenge

This document is aimed for contributors to the Challenge Series community
event which is held at GopherCon each year. The goal of this document is to
give a primer on the process of creating a new challenge for the Challenge
Series, and to follow through on balancing the difficulty of the challenge by
using hints.

## Background and Introduction

During the course of this document, we will explore the design of the "Holding
out for a Hero" challenge, from GopherCon 2024. We'll be starting from ideation
and proceeding through the end of the competition. An important thing to
remember regarding this document is that it is both a discussion of the sorts
of considerations that go into new challenges, as well as explanation of the
process for solving this challenge.

## Challenge Ideation

### Competition Theme

The ideation process begins with the theme we utilized for GopherCon 2024:

"Dial Up The Nostalgia"

To be more explicit, the GopherCon 2024 theme was the technology and media of
the period between approximately 1985 and 1995, a decade of technology to match
the decade of GopherCon.

Given this, we wanted our challenges to incorporate some aspects of 
contemporary technology from this time period. For myself, I have fond memories
of working on the Heathkit Educational Robot series, specifically the
[HERO ET-18](https://en.wikipedia.org/wiki/HERO_%28robot%29), though my time
with these devices came much later, in the later 90s and early 00s.

### What Makes A Possible Challenge?

For these devices, the memories that stick out in my mind involved programming
the three robots we had at my high school to give speeches, using the built-in
speech synthesizer. This spawned the idea of a challenge where the players were
required to, in some way, use the same speech synthesizer to retrieve a flag.

Some of the ideas I worked through in the stages of planning for this challenge
included:
- Acquiring a HERO ET-18 and bringing it to the conference. 
This was unfeasible due to cost and size of the device.
- Having the players reverse engineer the speech synthesizer and pass through
the speech code to get the flag from the other side.
While this is an interesting idea, this would involve a much larger time 
commitment than we really want the players to have to devote to a single
challenge.
- Building my own reverse engineered speech synthesizer and providing the 
players with a way to feed it instruction sets, forcing the players to filter
through multiple sets of instructions to find the right one, and using some
sort of metadata on the "right" file to be the flag.

In the end, the solution that I chose to pursue was inspired by different
aspects of the above ideas:

"Given a memory dump of a program from the HERO ET-18, find the Flag that has
been hidden within."

> **Note:** In reality, ET-18 programs were written from memory to a cassette tape,
and read back into memory from cassette as well. I elected to skip this
particular hardware requirement for solving the puzzle.

Now, with our problem statement set, we need to work backwards from a solution
(which we've yet to create!) toward the problem statement we've outlined above.

## Working Backwards

### Creating a Bounds to the Solution

In this case, I knew that I wanted the Flag to be encoded in
speech synthesizer code, and also, I wanted to be sure that the speech
synthesizer code wasn't the only code in the program. Given these constraints,
I started from the ET-18 Instruction Manual to get some information that I knew
existed, but did not have committed to memory from twenty years ago:

1. *What memory address can programs start at?* `003F`
2. *What memory address is the last available?* `0EE0`
3. *What are the specific command codes available to me?* Outlined in the manual
starting on page 25.

At this point, some additional context is needed for how the Heathkit ET-18
robot works. While the full details of the ET-18 are better found in the 
[manual](https://archive.org/details/heathkit-hero-1-ET-18), a summary of how
programs are created, stored, and run are useful context.

In short, each memory address in the robot's memory bank (`003F` -> `0EE0`)
stores a single byte. The value of that byte relates directly to a specific
instruction for the robot. Additionally, the robot is modal, and the robot will
switch modes based on these instructions as well. This capability is referred
to as "Robot Language" within the user manual, and while it is the sole method
we're considering for this challenge, it is not the only way of programming
these robots.

### Building Our Resources

Given all of the information above, we're able to create a reasonable facsimile
of the memory addresses and instructions of a small Heathkit ET-18 program,
which can be seen [here](./heathkit.mem). In this program, we give the robot a
number of instructions, primarily movement. However, at memory address `0047`,
we give the following instruction: `72 01 00`

> **Definition:** Phonemes are the smallest unit of spoken word in spoken
language, and can be used to construct any sounds when sequenced properly. The
ET-18 robot has specific codes for each of the standard phonemes, allowing it
to speak whatever we want.

This instruction tells the robot to start speaking the phoneme codes stored at
memory address location `0100` and proceed until it receives the "end speech"
directive, which is the value `3F`. We can see that this value exists at memory
address `0138`, meaning that the phoneme codes in range `0101` through `0137`
are spoken by the robot when this program is executed. 

While selecting the phonemes for specific word sounds and emphasis can take
some trial and error, in this case, being a best-effort-approximation is a
benefit, rather than a problem - Without the players needing to actually hear
the sounds being made by the robot, the sound descriptions outlined in the
manual can be "good enough" and still get the message across.

Given this, we've been able to successfully work backwards from our known
problem statement and the solution to the challenge, but now we need to move
forward toward the solution and validate that our players are going to be able to solve the challenge.

## Working Forwards

### Assumed Player Knowledge

For working through the problem, our players will start with just a small
amount of information - The problem statement we've outlined above.

We'll want to make a best-effort attempt to put aside the information we
already know, and try to think about solving the problem from the perspective
of someone who doesn't know the solution, or even the problem space as of yet.
It can actually be helpful to have a partner to work through this problem
alongside you and take notes on the places where they have to ask for help.
These places can (and should) be added as hints to the challenge.

In this case, the end users know only that they have the `heathkit.mem` file,
and that it's a memory representation of an ET-18 program. They can reason that
the user manual would contain more information, but things they won't know are:

1. What sort of things might be in the program.
2. That the ET-18 has a speech synthesizer.
3. Potentially, how phonemes work.

This means that, to start, we want to give hints for each of these. But within
the competition, unlocking additional hints will cost the team points, we also
give some small hints within the original message the players see. In this
case, we make a small riddle:

```
In the depths of the binary sea,
A robot speaks, then moves with glee.
Forward and back, then it halts,
A sequence of sounds, with no faults.
Listen close, decode the track,
The flag you seek is in its back.
```

This riddle hints to the player that there's something to do with a robot, that
there is some sort of movement and speech involved, and that finding the flag
requires listening or decoding some sort of audio.

Now that we've revised and settled on our initial challenge statement, we need
to ensure that our players are able to bridge the gap. 

### Building Hints

For reference, I've included a table here which outlines the hints and their
costs, but we'll dive more fully into them in this section.

| Hint # | Hint Description                                                                                     | Cost (Points) |
|--------|-----------------------------------------------------------------------------------------------------|---------------|
| 1      | Pay attention to the sequence of actions the robot performs. Identify the sections where it moves and where it speaks. The order of operations is crucial to understanding the hidden message.    | 1             |
| 2      | The file included provides a memory dump of a program for a Heathkit Educational Robot, model ET-18. 

There's multiple copies of the manual online, but you can also find one here:
https://archive.org/details/heathkit-hero-1-ET-18  | 5             |
| 3      | The instruction for "Begin Speaking" in the ET-18's Robot Language is "72", followed by a memory address where the speech program begins.    | 10            |

#### An Aside on Cost

> **Note**: The top three teams each year have so far received a variety of
fabulous prizes, which has been a big motivator for how these teams work on
challenges throughout the competition.

A balancing factor that we utilize for providing hints is that the hints must
be paid for with points that the team has already earned solving other
challenges during the competition. This means that the decision to purchase a
hint may put a team in a less competitive position if they're aiming for one of
the top three spots in the competition. We've settled on a system where,
generally, a hint will be of a small, medium, or large size, and in most
challenges, smaller hints will have to be purchased before the larger hints.

This means that if a team purchases a small (with a 1-point cost), a medium 
(with a 5-point cost), and a large (with a 10-point cost), they've spent 16
points to solve that challenge. Given that most challenges are worth 100
points, this is a pretty large investment to make, and we've found that this
balances out the risk tolerance of most of our teams.

#### Hint # 1

Now, in order to not make the challenge incredibly simple for the players, we
remove the original problem statement, and sprinkle that information throughout
the remaining hints. A good first hint for a challenge should be something that
is actually of varying quality based on the information the team has at hand.
For example, in this instance, our first hint is the following text:

```
Pay attention to the sequence of actions the robot performs. Identify the sections where it moves and where it speaks. The order of operations is crucial to understanding the hidden message.
```

For a team that has already found out that this is an ET-18 memory dump, 
they'll be able to understand that this means they need to trace the actions
being performed throughout the whole program to get the flag, and additionally
that the speech synthesizer is likely important. Given this information, they
have a competitive advantage against other teams that may need to buy more
hints in order to finish the challenge.

#### Hint # 2

However, for a team that hasn't figured that out yet, it's time to make it more
clear. Our second hint is as follows:

```
The file included provides a memory dump of a program for a Heathkit Educational Robot, model ET-18. 

There's multiple copies of the manual online, but you can also find one here:
https://archive.org/details/heathkit-hero-1-ET-18
```

Now, this allows us to get any team that needs a second hint onto fairly equal
footing with each other. They've got roughly the same exact information even if
they came by it naturally or had to buy the hint. But, some teams may still
need an extra hint, if they're short on time or they find the manual hard to
understand.

#### Hint # 3

For our third hint, we skip a lot of the work that has to be done to actually
research the robot and its programming, and we tell the player exactly what
instruction to be looking for:

```
The instruction for "Begin Speaking" in the ET-18's Robot Language is "72", followed by a memory address where the speech program begins.
```

This puts the player who has bought a third hint in the position where they can
quickly search the file for a `72`, check the memory address it references,
then proceed with checking those addresses for phonemes and finish solving the
puzzle. However, buying all three hints, as we mention above, represents a huge
opportunity cost.

## Conclusion

Hopefully, the explanation above has been helpful in starting to understand the
considerations and thought process that we typically put into crafting new
challenges for the Challenge Series. Additionally, I hope that it helps give 
you some ideas on where to start with crafting your own challenge ideas. If you
need to get an opinion on your challenge ideas, or if there's anything you
want clarification on, don't hesitate to reach out in the #gophercon-cs-admin
Slack channel, or to either of us directly.