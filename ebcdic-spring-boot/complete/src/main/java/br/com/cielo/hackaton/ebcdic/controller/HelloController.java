package br.com.cielo.hackaton.ebcdic.controller;

import java.math.BigDecimal;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import br.com.cielo.hackaton.ebcdic.domain.SettlementMovement;
import br.com.cielo.hackaton.ebcdic.persistence.SettlementMovementRepository;

@RestController
public class HelloController {

	@Autowired
	SettlementMovementRepository repository;

	@RequestMapping("/test")
	public String home() {
		return "Iniciando servico Spring Boot";
	}

	@RequestMapping(value = "/ajustes", method = RequestMethod.POST)
	public HttpStatus create(@RequestBody SettlementMovement movement) throws SQLException {
		repository.insertMovement(getMocks());

		return HttpStatus.OK;
	}

	@RequestMapping(value = "/ajustesMock")
	public HttpStatus create() throws SQLException {
		repository.insertMovement(getMocks());

		return HttpStatus.OK;
	}
	
	@RequestMapping(value = "/ajustesAllMock")
	public HttpStatus createAllMocs() throws SQLException {
		repository.insertMovementBatchCommand(getAllMocks());

		return HttpStatus.OK;
	}

	private List<SettlementMovement> getAllMocks() {

		ArrayList<SettlementMovement> movements = new ArrayList<>();

		for (int i = 0; i < 1000000; i++) {

			movements.add(getMocks());
		}

		return movements;
	}

	private SettlementMovement getMocks() {

		SettlementMovement movementMock = new SettlementMovement(new Date(), new Integer("40"), 1003299951L, new Date(),
				new BigDecimal(100), new BigDecimal(120), new BigDecimal(20), 888889L, "986", "", new Date(),
				new Integer("10"), 22, UUID.randomUUID().toString(), Integer.valueOf(2));
		return movementMock;
	}
}