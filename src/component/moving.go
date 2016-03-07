package component

import ()

const MASK_MOVING = COMPONENT_POSITION | COMPONENT_VELOCITY | COMPONENT_BOX

func UpdateMoving(c *Collection) {
  for i, mask := range c.Mask {
    if (mask & MASK_MOVING) != MASK_MOVING {
      continue
    }

    pos := &c.Position[i]
    vel := &c.Velocity[i]
    intersections := &c.Intersections[i]
    box := c.Box[i]

    box.X += pos.X
    box.Y += pos.Y

    intersections.X = 0
    intersections.Y = 0


    box.X += vel.X

    for _, wall := range c.Walls {
    	if wall.W == 0 || wall.H == 0 {
    		continue
    	}

    	if box.Intersects(&wall) {
    		if vel.X > 0 {
    			box.X = wall.X - box.W
    			intersections.X = 1
    		} else {
    			intersections.X = -1
    			box.X = wall.X + wall.W
    		}
    	}
    }

    box.Y += vel.Y

    for _, wall := range c.Walls {
    	if wall.W == 0 || wall.H == 0 {
    		continue
    	}

    	if box.Intersects(&wall) {
    		if vel.Y > 0 {
    			box.Y = wall.Y - box.H
    			intersections.Y = 1
    		} else {
    			box.Y = wall.Y + wall.H
    			intersections.Y = -1
    		}
    	}
    }


    pos.X = box.X
    pos.Y = box.Y


    vel.X *= 0.8
    vel.Y += 0.3
  }
}
