package enum

// iota: Essa keyword faz com que o Go atribua o valor 0 à primeira constante e, de forma sequencial, vá incrementando em 1 o valor atribuído a cada constante.

type Status uint8

const (
	Created Status = iota
	Pending
	Approved
	Rejected
	Accepted
	Test
)

func (s Status) String() string {
	switch s {
	case Created:
		return "created"
	case Pending:
		return "pending"
	case Approved:
		return "approved"
	case Rejected:
		return "rejected"
	case Accepted:
		return "accepted"
	}
	return "unknown"
}
