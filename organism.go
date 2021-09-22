package organism

// New create new organism
func New() *Organism {
	o := &Organism{}
	o.core = o.GrowLimb()
	return o
}

// Organism is your application, it has core and limbs. Each Limb describes important part of your app, and if this part
// is dead, then organism is partially dead too (liveness probe). If each Limb of Organism, and the core too, are ready
// to work, then whole Organism ready too work too (readiness probe)
type Organism struct {
	limbs []*Limb
	core  *Limb
}

// GrowLimb grow new Limb for Organism and returns it
func (o *Organism) GrowLimb() *Limb {
	limb := newLimb()
	o.limbs = append(o.limbs, limb)

	return limb
}

func (o *Organism) IsReady() bool {
	for _, l := range o.limbs {
		if !l.IsReady() {
			return false
		}
	}

	return true
}

func (o *Organism) IsAlive() bool {
	for _, l := range o.limbs {
		if !l.IsAlive() {
			return false
		}
	}

	return true
}

// Ready marks core of Organism as ready
func (o *Organism) Ready() {
	o.core.Ready()
}

// Ready marks core of Organism as dead
func (o *Organism) Die() {
	o.core.Die()
}

func newLimb() *Limb {
	return &Limb{isAlive: true}
}

type Limb struct {
	isReady bool
	isAlive bool
}

// Ready marks Limb as ready
func (l *Limb) Ready() {
	l.isReady = true
}

// Die marks Limb as dead (not alive)
func (l *Limb) Die() {
	l.isAlive = false
}

func (l *Limb) IsReady() bool {
	return l.isReady
}

func (l *Limb) IsAlive() bool {
	return l.isAlive
}
