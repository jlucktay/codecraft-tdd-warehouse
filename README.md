# Exercise 3 - CD Warehouse

## _Identify the user actions we need to test..._

- Customer
  - buy a CD (hottest path)
    - includes paying for a CD through an external payment provider
      - would tie into stock level; reduce count by number purchased
    - leave a review after buying
      - numbered 1 through 10
      - some (optional) review text as well
  - search by title
  - search by artist
- RecordLabel
  - send a batch of CDs to us at the warehouse
- Warehouse
  - keep stock count of each CD title in the warehouse

## _... and make a Test List for each of them to drive your TDD_

- sell a single CD to a customer and take payment ✅
- attempt to sell a CD that doesn't exist ✅

- create a definition of a CD inside our warehouse
  - includes GUID, price, title, artist, label, stock level
  - reviews field/sub-type
    - average score, optional body text
- batch receipt of CD stock from record label into warehouse
- look up a CD by its title
- look up a CD by its artist
- look up a CD by both title and artist
- sell multiple copies of the same CD to a customer and take payment
- sell multiple different CDs to a customer and take payment
- sell a CD to a customer, take payment, and their review score
- sell a CD to a customer, take payment, their review score, and review text
