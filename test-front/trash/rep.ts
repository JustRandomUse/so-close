generateReport() {
    this.inventoryService.getReport(this.selectedCabinet).subscribe(reportData => {
      this.report = reportData;
    });
  }