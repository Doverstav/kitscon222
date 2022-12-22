# KitsCon 22.2 Demo

This repo contains the exploratory results of looking into Charm/Bubbletea and Cobra in preparation for my KitsCon 22.2 presentation, along with the end results of the live coding done during the presentation and the demoscript and presentation.

In `/charm_demo` and `/cobra_demo` are the playgrounds where I tested out Cobra and Charm/Bubbletea.

In `/livecoding` are the results of the live coding done during the presentation.

In `/docs` is the demo script which lays out how to do the live coding and the slides (in swedish) that were shown during the presentation. 

Data is stored using [badgerDB](https://github.com/dgraph-io/badger).

Database "design":
[kitscons]: Array of ids of all kitscons
[UUID]: Object containing all information about a kitscon (name, description, id, presentations)
[UUID]: Object containing all information about a presentation (title, presenter, rating, review, id)

## Test apps

The general idea behind the test app is a basic CRUD app that allows the user to keep track of presentation seen during KitsCons and what they though of them. So the user should be able to add: 
- KitsCons - Which will have a name, a description and a list of presentation ids (in addition to an id)
- Presentations - Which are related to a Kitscon and have a name, presenter, description, rating and a review (in addition to an id)

I decided to store this data locally on disk using [badgerDB](https://github.com/dgraph-io/badger). Database "design" looks like this:  
- `[kitscons]`: Array of ids of all kitscons  
- `[UUID]`: Object containing all information about a kitscon (name, description, id, presentations)  
- `[UUID]`: Object containing all information about a presentation (title, presenter, rating, review, id)  

### Cobra demo

#### Commands
- `kitscon add conf "<NAME>" "<DESCRIPTION>"` -> Adds new conference with name and description  
- `kitscon add presentation "<NAME>" "<PRESENTER>" "<DESCRIPTION>" "<RATING>" "<REVIEW>" --conf="<CONFERENCE NAME>"` -> Adds new presentation with name and description under conference  
- `kitscon list` -> List all conferences, presentations and reviews  
- `kitscon list confs` -> List conferences  
- `kitscon list presentations --conf="<CONFERENCE NAME>` -> List presentation under conference  

### Charm demo

This charm demo consists of a few screens which are described below.

#### KITSCON_LIST:
Initial view where you see a list of all added KitsCons. From here you can go to `ADD_NEW_KITSCON` by pressing some hotkey or button, or to `PRESENTATIONS_LIST` by selecting a conference from the list.

#### ADD_NEW_KITSCON:
In this view you can add a new KitsCon with a name/title and a description.

#### PRESENTATIONS_LIST:
Shows a list of all presentations added under a conference. If a review has been submitted, show rating under presentation name. From here you can go to `ADD_NEW_PRESENTATION` by pressing some hotkey or button, or to `PRESENTATION` by selecting a presentation from the list

#### ADD_NEW_PRESENTATION:
Allows the user to add a new presentation under a conference. User can add a presentations name/title, who presented, a description, a rating and some thoughts on the presentation.

#### PRESENTATION:
Shows details about a presentations, i.e. the info added in `ADD_NEW_PRESENTATION`