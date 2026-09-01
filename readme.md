# Project Matrix

## The Problem

After having planed a spec for a feature, major refactoring, bug fixes etc., it's time for implementation. The now traditional way of doing it
runs like this:

1. You prompt your coding agent:
```
Implement spec X with issue id #3
```
2. Depending on the size of the task you wait half an hour to several hours. Cannot close laptop. Must wait.
3. After the implementation you review the changes. Either with another agent or manually. But initiated manually either way.
4. Make needed changes, again either with an agent or manually
5. Merge changes

There are a couple of problems with this approach:
1. Too much for a single agent. You are way over the smart zone (or how I call it: the non-dumb-zone) of the context, which is briefly
the interval from 0 tokens to 100K tokens.
2. No rules for implementation. You are just telling the agent to "implement". If your existing code is garbage, you get garbage out. No way of preventing it.
3. Exhausting review with multiple unnecessary cycles. If the agent touched 200 files, you cannot possibly review it efficiently.
4. Unnecessary or missing tests. The agent won't use TDD by default.
5. You wasted tokens and effectively money. Better off implementing it "manually".

### A couple of improvements

1. You use some flavour of TDD skill to implement. Huge improvement in code quality.
2. You use a fresh context window for each sub issue, rather than the big chunk of PRD. This way you stay in the non-dumb-zone more often.
3. You cut large pieces to smaller tickets. Huge improvement for code review and code quality and actually hitting the ACs. 
4. You use some kind of Ralph Loop implementations to work on tickets. Less time in front of keyboard.
More AFK time. Still laptop must stay on. Cannot share progress with team or monitor and manage it remotely. 

## Our Solution

## What Matrix Is

## What Matrix Is Not

## Setup

## Usage

## Contribution